package entity

import (
	"time"

	"gorm.io/gorm"
)

// ============================================================
// ENUM Types
// ============================================================

type UserRole string
type OrderStatus string
type PaymentStatus string

const (
	RoleAdmin    UserRole = "admin"
	RoleSupplier UserRole = "supplier"
	RoleSPPG     UserRole = "sppg"
)

const (
	OrderPending    OrderStatus = "pending"
	OrderProcessing OrderStatus = "processing"
	OrderShipped    OrderStatus = "shipped"
	OrderCompleted  OrderStatus = "completed"
	OrderCancelled  OrderStatus = "cancelled"
)

const (
	PaymentUnpaid              PaymentStatus = "unpaid"
	PaymentWaitingConfirmation PaymentStatus = "waiting_confirmation"
	PaymentPaid                PaymentStatus = "paid"
	PaymentFailed              PaymentStatus = "failed"
)

// ============================================================
// Base Model
// ============================================================

type Base struct {
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"          json:"-"`
}