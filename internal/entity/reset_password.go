package entity

import (
	"time"

	"github.com/google/uuid"
)

type ResetPassword struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"                              json:"user_id"`
	Token     string    `gorm:"type:varchar(255);not null;uniqueIndex"          json:"token"`
	ExpiredAt time.Time `gorm:"not null"                                        json:"expired_at"`
	IsUsed    bool      `gorm:"not null;default:false"                          json:"is_used"`
	CreatedAt time.Time `gorm:"autoCreateTime"                                  json:"created_at"`
}

func (ResetPassword) TableName() string { return "reset_passwords" }