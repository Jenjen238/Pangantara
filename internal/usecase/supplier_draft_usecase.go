package usecase

import (
	"errors"
	"sppg-backend/internal/entity"
	"sppg-backend/internal/model"
	"sppg-backend/internal/repository"

	"github.com/google/uuid"
)

func SaveDraft(req model.SaveDraftRequest) (*entity.SupplierDraft, error) {
	draft := &entity.SupplierDraft{
		ID:            uuid.New(),
		UserID:        req.UserID,
		StoreName:     req.StoreName,
		Address:       req.Address,
		ContactNumber: req.ContactNumber,
		Category:      req.Category,
		SourceType:    req.SourceType,
		BusinessDesc:  req.BusinessDesc,
		CurrentStep:   req.CurrentStep,
		Status:        entity.DraftStatusDraft,
	}

	if err := repository.CreateOrUpdateDraft(draft); err != nil {
		return nil, err
	}

	// Ambil data terbaru setelah disimpan
	saved, err := repository.GetDraftByUserID(req.UserID)
	if err != nil {
		return nil, err
	}

	return saved, nil
}

func GetDraft(userID uuid.UUID) (*entity.SupplierDraft, error) {
	draft, err := repository.GetDraftByUserID(userID)
	if err != nil {
		return nil, errors.New("draft tidak ditemukan")
	}
	return draft, nil
}

func UpdateDraftDocument(userID uuid.UUID, docType, path string) error {
	return repository.UpdateDraftDocument(userID, docType, path)
}

func SubmitDraft(req model.SubmitDraftRequest) (*entity.Supplier, error) {
	// Ambil draft
	draft, err := repository.GetDraftByUserID(req.UserID)
	if err != nil {
		return nil, errors.New("draft tidak ditemukan")
	}

	// Validasi field wajib
	if draft.StoreName == nil || *draft.StoreName == "" {
		return nil, errors.New("nama toko wajib diisi")
	}

	// Buat supplier dari draft
	supplier := &entity.Supplier{
		SupplierID:         uuid.New(),
		UserID:             draft.UserID,
		StoreName:          *draft.StoreName,
		Address:            draft.Address,
		ContactNumber:      draft.ContactNumber,
		Category:           draft.Category,
		SourceType:         draft.SourceType,
		BusinessDesc:       draft.BusinessDesc,
		NIBDocument:        draft.NIBDocument,
		HalalDocument:      draft.HalalDocument,
		OtherDocument:      draft.OtherDocument,
		VerificationStatus: entity.VerificationPending,
	}

	if err := repository.CreateSupplier(supplier); err != nil {
		return nil, err
	}

	// Update status draft jadi submitted
	if err := repository.SubmitDraft(req.UserID); err != nil {
		return nil, err
	}

	return supplier, nil
}

func DeleteDraft(userID uuid.UUID) error {
	return repository.DeleteDraft(userID)
}