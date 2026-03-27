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
		return errors.New("transaksi tidak ditemukan")
	}

	// Update payment status
	if err := repository.UpdateTransactionStatus(id, entity.PaymentStatus(req.PaymentStatus)); err != nil {
		return err
	}

	// Sync order status berdasarkan payment status
	if req.PaymentStatus == string(entity.PaymentPaid) {
		_ = repository.UpdateOrderStatus(transaction.OrderID, entity.OrderProcessing)
	} else if req.PaymentStatus == string(entity.PaymentFailed) {
		_ = repository.UpdateOrderStatus(transaction.OrderID, entity.OrderCancelled)
	}

	// Kirim email notifikasi
	if req.PaymentStatus == string(entity.PaymentPaid) || req.PaymentStatus == string(entity.PaymentFailed) {
		order, err := repository.GetOrderByID(transaction.OrderID)
		if err != nil {
			log.Println("[EMAIL ERROR] GetOrderByID gagal:", err)
			return nil
		}

		sppg, err := repository.GetSPPGByID(order.SPPGID)
		if err != nil {
			log.Println("[EMAIL ERROR] GetSPPGByID gagal:", err)
			return nil
		}

		user, err := repository.GetUserByID(sppg.UserID)
		if err != nil {
			log.Println("[EMAIL ERROR] GetUserByID gagal:", err)
			return nil
		}

		orderIDStr := transaction.OrderID.String()
		log.Println("[EMAIL] Mengirim email ke:", user.Email)

		go func() {
			var emailErr error
			if req.PaymentStatus == string(entity.PaymentPaid) {
				emailErr = email.SendPaymentConfirmedEmail(user.Email, user.Name, orderIDStr, order.TotalAmount)
			} else {
				emailErr = email.SendPaymentRejectedEmail(user.Email, user.Name, orderIDStr, order.TotalAmount)
			}
			if emailErr != nil {
				log.Println("[EMAIL ERROR] Gagal kirim email:", emailErr)
			} else {
				log.Println("[EMAIL SUCCESS] Email terkirim ke:", user.Email)
			}
		}()
	}

	return nil
}

func DeleteTransaction(id uuid.UUID) error {
	return repository.DeleteTransaction(id)
}