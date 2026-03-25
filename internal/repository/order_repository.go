package repository

import (
	"sppg-backend/internal/entity"
	"sppg-backend/pkg/postgres"

	"github.com/google/uuid"
)

func CreateOrder(o *entity.Order) error {
	return postgres.DB.Create(o).Error
}

func GetAllOrder() ([]entity.Order, error) {
	var list []entity.Order
	err := postgres.DB.Order("order_date DESC").Find(&list).Error
	return list, err
}

func GetOrderByID(id uuid.UUID) (*entity.Order, error) {
	var o entity.Order
	err := postgres.DB.Preload("OrderDetail").Preload("Transaction").
		First(&o, "order_id = ?", id).Error
	return &o, err
}

func GetOrderBySPPGID(sppgID uuid.UUID) ([]entity.Order, error) {
	var list []entity.Order
	err := postgres.DB.Where("sppg_id = ?", sppgID).
		Order("order_date DESC").Find(&list).Error
	return list, err
}

func GetOrderByStatus(status entity.OrderStatus) ([]entity.Order, error) {
	var list []entity.Order
	err := postgres.DB.Where("order_status = ?", status).
		Order("order_date DESC").Find(&list).Error
	return list, err
}

func UpdateOrderStatus(id uuid.UUID, status entity.OrderStatus) error {
	return postgres.DB.Model(&entity.Order{}).
		Where("order_id = ?", id).
		Update("order_status", status).Error
}

func DeleteOrder(id uuid.UUID) error {
	return postgres.DB.Delete(&entity.Order{}, "order_id = ?", id).Error
}

func CountOrder() (int64, error) {
	var count int64
	err := postgres.DB.Model(&entity.Order{}).Count(&count).Error
	return count, err
}

func CountOrderByStatus(status entity.OrderStatus) (int64, error) {
	var count int64
	err := postgres.DB.Model(&entity.Order{}).Where("order_status = ?", status).Count(&count).Error
	return count, err
}

func GetOrdersFiltered(status entity.OrderStatus, sppgID *uuid.UUID, startDate, endDate *string, page, limit int) ([]entity.Order, int64, error) {
	var list []entity.Order
	var total int64

	db := postgres.DB.Model(&entity.Order{})

	if status != "" {
		db = db.Where("order_status = ?", status)
	}
	if sppgID != nil {
		db = db.Where("sppg_id = ?", sppgID)
	}
	if startDate != nil && *startDate != "" {
		db = db.Where("order_date >= ?", *startDate)
	}
	if endDate != nil && *endDate != "" {
		db = db.Where("order_date <= ?", *endDate+" 23:59:59")
	}

	db.Count(&total)

	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	err := db.Order("order_date DESC").Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}