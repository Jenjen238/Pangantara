package usecase

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"sppg-backend/config"
	"sppg-backend/internal/entity"
	"sppg-backend/internal/model"
	"sppg-backend/internal/repository"
	"sppg-backend/pkg/email"
	"sppg-backend/pkg/payment"
	"strconv"

	"github.com/google/uuid"
)

func CreatePayment(orderID string) (*payment.PaymentResponse, error) {
	// Ambil data order
	id, err := uuid.Parse(orderID)
	if err != nil {
		return nil, errors.New("order ID tidak valid")
	}

	order, err := repository.GetOrderByID(id)
	if err != nil {
		return nil, errors.New("order tidak ditemukan")
	}

	// Ambil data SPPG
	sppg, err := repository.GetSPPGByID(order.SPPGID)
	if err != nil {
		return nil, errors.New("SPPG tidak ditemukan")
	}

	// Ambil data user
	user, err := repository.GetUserByID(sppg.UserID)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	// Ambil detail order
	details, err := repository.GetOrderDetailByOrderID(order.OrderID)
	if err != nil {
		return nil, errors.New("detail order tidak ditemukan")
	}

	// Setup items
	var items []payment.PaymentItem
	for _, detail := range details {
		product, err := repository.GetProductByID(detail.ProductID)
		if err != nil {
			continue
		}
		items = append(items, payment.PaymentItem{
			ID:       detail.ProductID.String(),
			Name:     product.ProductName,
			Price:    int64(product.Price),
			Quantity: detail.Quantity,
		})
	}

	// Buat payment di Midtrans
	paymentReq := payment.PaymentRequest{
		OrderID:       orderID,
		Amount:        int64(order.TotalAmount),
		CustomerName:  user.Name,
		CustomerEmail: user.Email,
		Items:         items,
	}

	resp, err := payment.CreatePayment(paymentReq)
	if err != nil {
		return nil, err
	}

	// Buat transaksi di database
	transaction := &entity.Transaction{
		TransactionID: uuid.New(),
		OrderID:       order.OrderID,
		PaymentMethod: nil,
		PaymentStatus: entity.PaymentUnpaid,
		AmountPaid:    order.TotalAmount,
	}
	repository.CreateTransaction(transaction)

	return resp, nil
}

func HandleMidtransNotification(notif model.MidtransNotification) error {
	// Verifikasi signature key
	signatureInput := fmt.Sprintf("%s%s%s%s",
		notif.OrderID,
		"200",
		notif.GrossAmount,
		config.AppConfig.MidtransServerKey,
	)
	hash := sha512.New()
	hash.Write([]byte(signatureInput))
	expectedSignature := fmt.Sprintf("%x", hash.Sum(nil))

	if expectedSignature != notif.SignatureKey {
		return errors.New("signature tidak valid")
	}

	// Parse order ID
	orderID, err := uuid.Parse(notif.OrderID)
	if err != nil {
		return errors.New("order ID tidak valid")
	}

	// Update status berdasarkan notifikasi Midtrans
	var paymentStatus entity.PaymentStatus
	var orderStatus entity.OrderStatus

	switch notif.TransactionStatus {
	case "capture", "settlement":
		if notif.FraudStatus == "accept" || notif.FraudStatus == "" {
			paymentStatus = entity.PaymentPaid
			orderStatus = entity.OrderProcessing
		}
	case "pending":
		paymentStatus = entity.PaymentWaitingConfirmation
		orderStatus = entity.OrderPending
	case "deny", "cancel", "expire":
		paymentStatus = entity.PaymentFailed
		orderStatus = entity.OrderCancelled
	default:
		return nil
	}

	// Ambil transaksi by order ID
	transaction, err := repository.GetTransactionByOrderID(orderID)
	if err != nil {
		return errors.New("transaksi tidak ditemukan")
	}

	// Update payment status
	grossAmount, _ := strconv.ParseFloat(notif.GrossAmount, 64)
	repository.UpdateTransactionStatus(transaction.TransactionID, paymentStatus)
	repository.UpdateOrderStatus(orderID, orderStatus)

	// Kirim email notifikasi
	order, err := repository.GetOrderByID(orderID)
	if err != nil {
		return nil
	}

	sppg, err := repository.GetSPPGByID(order.SPPGID)
	if err != nil {
		return nil
	}

	user, err := repository.GetUserByID(sppg.UserID)
	if err != nil {
		return nil
	}

	orderIDStr := orderID.String()
	if paymentStatus == entity.PaymentPaid {
		go email.SendPaymentConfirmedEmail(user.Email, user.Name, orderIDStr, grossAmount)
	} else if paymentStatus == entity.PaymentFailed {
		go email.SendPaymentRejectedEmail(user.Email, user.Name, orderIDStr, grossAmount)
	}

	return nil
}