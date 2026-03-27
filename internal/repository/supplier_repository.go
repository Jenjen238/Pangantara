package repository

import (
	"sppg-backend/internal/entity"
	"sppg-backend/pkg/postgres"

	"github.com/google/uuid"
)

func CreateSupplier(s *entity.Supplier) error {
	return postgres.DB.Create(s).Error
}

func GetAllSupplier() ([]entity.Supplier, error) {
	var list []entity.Supplier
	err := postgres.DB.Order("store_name ASC").Find(&list).Error
	return list, err
}

func GetSupplierByVerificationStatus(status entity.VerificationStatus) ([]entity.Supplier, error) {
	var list []entity.Supplier
	err := postgres.DB.Where("verification_status = ?", status).Order("created_at DESC").Find(&list).Error
	return list, err
}

func GetSupplierByID(id uuid.UUID) (*entity.Supplier, error) {
	var s entity.Supplier
	err := postgres.DB.First(&s, "supplier_id = ?", id).Error
	return &s, err
}

func GetSupplierByUserID(userID uuid.UUID) (*entity.Supplier, error) {
	var s entity.Supplier
	err := postgres.DB.Where("user_id = ?", userID).First(&s).Error
	return &s, err
}

func UpdateSupplier(id uuid.UUID, data map[string]interface{}) error {
	return postgres.DB.Model(&entity.Supplier{}).Where("supplier_id = ?", id).Updates(data).Error
}

func VerifySupplier(id uuid.UUID, status entity.VerificationStatus, notes *string) error {
	data := map[string]interface{}{
		"verification_status": status,
	}
	if notes != nil {
		data["admin_notes"] = notes
	}
	return postgres.DB.Model(&entity.Supplier{}).Where("supplier_id = ?", id).Updates(data).Error
}

func DeleteSupplier(id uuid.UUID) error {
	return postgres.DB.Delete(&entity.Supplier{}, "supplier_id = ?", id).Error
}

func CountSupplier() (int64, error) {
	var count int64
	err := postgres.DB.Model(&entity.Supplier{}).Count(&count).Error
	return count, err
}

func CountSupplierByStatus(status entity.VerificationStatus) (int64, error) {
	var count int64
	err := postgres.DB.Model(&entity.Supplier{}).Where("verification_status = ?", status).Count(&count).Error
	return count, err
}

func SearchSupplier(keyword, category, status string) ([]entity.Supplier, error) {
	var list []entity.Supplier
	db := postgres.DB.Model(&entity.Supplier{})

	if keyword != "" {
		db = db.Where("store_name ILIKE ?", "%"+keyword+"%")
	}
	if category != "" {
		db = db.Where("category = ?", category)
	}
	if status != "" {
		db = db.Where("verification_status = ?", status)
	}

	err := db.Order("store_name ASC").Find(&list).Error
	return list, err
}