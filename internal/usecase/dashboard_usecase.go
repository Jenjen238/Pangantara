package usecase

import (
	"sppg-backend/internal/entity"
	"sppg-backend/internal/model"
	"sppg-backend/internal/repository"
)

func GetDashboardSummary() (*model.DashboardSummary, error) {
	totalSupplier, err := repository.CountSupplier()
	if err != nil {
		return nil, err
	}
	supplierPending, _ := repository.CountSupplierByStatus(entity.VerificationPending)
	supplierApproved, _ := repository.CountSupplierByStatus(entity.VerificationApproved)
	supplierRejected, _ := repository.CountSupplierByStatus(entity.VerificationRejected)

	totalSPPG, _ := repository.CountSPPG()

	totalOrder, _ := repository.CountOrder()
	orderPending, _ := repository.CountOrderByStatus(entity.OrderPending)
	orderProcessing, _ := repository.CountOrderByStatus(entity.OrderProcessing)
	orderShipped, _ := repository.CountOrderByStatus(entity.OrderShipped)
	orderCompleted, _ := repository.CountOrderByStatus(entity.OrderCompleted)
	orderCancelled, _ := repository.CountOrderByStatus(entity.OrderCancelled)

	return &model.DashboardSummary{
		TotalSupplier:    totalSupplier,
		SupplierPending:  supplierPending,
		SupplierApproved: supplierApproved,
		SupplierRejected: supplierRejected,
		TotalSPPG:        totalSPPG,
		TotalOrder:       totalOrder,
		OrderPending:     orderPending,
		OrderProcessing:  orderProcessing,
		OrderShipped:     orderShipped,
		OrderCompleted:   orderCompleted,
		OrderCancelled:   orderCancelled,
	}, nil
}