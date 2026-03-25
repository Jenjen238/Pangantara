package repository

import (
	"sppg-backend/internal/entity"
	"sppg-backend/pkg/postgres"

	"github.com/google/uuid"
)

func CreateOrUpdateDraft(draft *entity.SupplierDraft) error {
	existing, err := GetDraftByUserID(draft.UserID)
	if err != nil {
		// Draft belum ada, buat baru
		return postgres.DB.Create(draft).Error
	}
	// Draft sudah ada, update
	return postgres.DB.Model(existing).Updates(map[string]interface{}{
		"store_name":     draft.StoreName,
		"address":        draft.Address,
		"contact_number": draft.ContactNumber,
		"category":       draft.Category,
		"source_type":    draft.SourceType,
		"business_desc":  draft.BusinessDesc,
		"current_step":   draft.CurrentStep,
	}).Error
}

func GetDraftByUserID(userID uuid.UUID) (*entity.SupplierDraft, error) {
	var draft entity.SupplierDraft
	err := postgres.DB.Where("user_id = ?", userID).First(&draft).Error
	return &draft, err
}

func UpdateDraftDocument(userID uuid.UUID, docType string, path string) error {
	data := map[string]interface{}{}
	switch docType {
	case "nib":
		data["nib_document"] = path
	case "halal":
		data["halal_document"] = path
	case "other":
		data["other_document"] = path
	}
	return postgres.DB.Model(&entity.SupplierDraft{}).
		Where("user_id = ?", userID).
		Updates(data).Error
}

func SubmitDraft(userID uuid.UUID) error {
	return postgres.DB.Model(&entity.SupplierDraft{}).
		Where("user_id = ?", userID).
		Update("status", entity.DraftStatusSubmitted).Error
}

func DeleteDraft(userID uuid.UUID) error {
	return postgres.DB.Where("user_id = ?", userID).
		Delete(&entity.SupplierDraft{}).Error
}