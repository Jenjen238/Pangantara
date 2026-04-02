package usecase

import (
	"errors"
	"log"
	"sppg-backend/internal/entity"
	"sppg-backend/internal/model"
	"sppg-backend/internal/repository"
	"sppg-backend/pkg/email"
	"time"

	"github.com/google/uuid"
)

func CreateTransaction(req model.CreateTransactionRequest) (*entity.Transaction, error) {
	now := time.Now()
	transaction := &entity.Transaction{
		TransactionID: uuid.New(),
		OrderID:       req.OrderID,
		PaymentMethod: req.PaymentMethod,
		PaymentStatus: entity.PaymentWaitingConfirmation,
		PaymentProof:  req.PaymentProof,
		PaymentDate:   &now,
		AmountPaid:    req.AmountPaid,
	}
	return transaction, repository.CreateTransaction(transaction)
}

func GetAllTransaction() ([]entity.Transaction, error) {
	return repository.GetAllTransaction()
}

func GetTransactionByID(id uuid.UUID) (*entity.Transaction, error) {
	return repository.GetTransactionByID(id)
}

func GetTransactionByOrderID(orderID uuid.UUID) (*entity.Transaction, error) {
	return repository.GetTransactionByOrderID(orderID)
}

func UpdateTransactionStatus(id uuid.UUID, req model.UpdatePaymentStatusRequest) error {
	transaction, err := repository.GetTransactionByID(id)
	if err != nil {
		return errors.New("transaction not found")
	}

	// Tentukan order status berdasarkan payment status
	var orderStatus entity.OrderStatus
	switch entity.PaymentStatus(req.PaymentStatus) {
	case entity.PaymentPaid:
		orderStatus = entity.OrderProcessing
	case entity.PaymentFailed:
		orderStatus = entity.OrderCancelled
	default:
		return repository.UpdateTransactionStatus(id, entity.PaymentStatus(req.PaymentStatus))
	}

	if err := repository.UpdateTransactionAndOrderStatus(
		id,
		entity.PaymentStatus(req.PaymentStatus),
		transaction.OrderID,
		orderStatus,
	); err != nil {
		return errors.New("failed to update payment and order status")
	}

	// Kirim email notifikasi
	order, err := repository.GetOrderByID(transaction.OrderID)
	if err != nil {
		log.Println("[EMAIL ERROR] GetOrderByID failed:", err)
		return nil
	}

	sppg, err := repository.GetSPPGByID(order.SPPGID)
	if err != nil {
		log.Println("[EMAIL ERROR] GetSPPGByID failed:", err)
		return nil
	}

	user, err := repository.GetUserByID(sppg.UserID)
	if err != nil {
		log.Println("[EMAIL ERROR] GetUserByID failed:", err)
		return nil
	}

	orderIDStr := transaction.OrderID.String()
	go func() {
		var emailErr error
		if req.PaymentStatus == string(entity.PaymentPaid) {
			emailErr = email.SendPaymentConfirmedEmail(user.Email, user.Name, orderIDStr, order.TotalAmount)
		} else {
			emailErr = email.SendPaymentRejectedEmail(user.Email, user.Name, orderIDStr, order.TotalAmount)
		}
		if emailErr != nil {
			log.Println("[EMAIL ERROR] Failed to send email:", emailErr)
		} else {
			log.Println("[EMAIL SUCCESS] Email sent to:", user.Email)
		}
	}()

	return nil
}

func DeleteTransaction(id uuid.UUID) error {
	return repository.DeleteTransaction(id)
}
