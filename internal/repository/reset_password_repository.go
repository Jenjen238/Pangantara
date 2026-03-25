package repository

import (
	"sppg-backend/internal/entity"
	"sppg-backend/pkg/postgres"
	"time"

	"github.com/google/uuid"
)

func CreateResetPassword(r *entity.ResetPassword) error {
	return postgres.DB.Create(r).Error
}

func GetResetPasswordByToken(token string) (*entity.ResetPassword, error) {
	var r entity.ResetPassword
	err := postgres.DB.Where("token = ? AND is_used = false AND expired_at > ?", token, time.Now()).
		First(&r).Error
	return &r, err
}

func MarkResetPasswordAsUsed(id uuid.UUID) error {
	return postgres.DB.Model(&entity.ResetPassword{}).
		Where("id = ?", id).
		Update("is_used", true).Error
}

func DeleteExpiredResetPassword(userID uuid.UUID) error {
	return postgres.DB.Where("user_id = ? AND expired_at < ?", userID, time.Now()).
		Delete(&entity.ResetPassword{}).Error
}