package usecase

import (
	"errors"
	"sppg-backend/internal/entity"
	"sppg-backend/internal/model"
	"sppg-backend/internal/repository"
	"sppg-backend/pkg/email"

	"github.com/google/uuid"
)

func CreateSupplier(req model.CreateSupplierRequest) (*entity.Supplier, error) {
	supplier := &entity.Supplier{
		SupplierID:         uuid.New(),
		UserID:             req.UserID,
		StoreName:          req.StoreName,
		Address:            req.Address,
		ContactNumber:      req.ContactNumber,
		Category:           req.Category,
		SourceType:         req.SourceType,
		BusinessDesc:       req.BusinessDesc,
		NIBDocument:        req.NIBDocument,
		HalalDocument:      req.HalalDocument,
		OtherDocument:      req.OtherDocument,
		AdminNotes:         req.AdminNotes,
		VerificationStatus: entity.VerificationPending,
	}
	return supplier, repository.CreateSupplier(supplier)
}

func GetAllSupplier() ([]entity.Supplier, error) {
	return repository.GetAllSupplier()
}

func GetSupplierByVerificationStatus(status string) ([]entity.Supplier, error) {
	return repository.GetSupplierByVerificationStatus(entity.VerificationStatus(status))
}

func GetSupplierByID(id uuid.UUID) (*entity.Supplier, error) {
	return repository.GetSupplierByID(id)
}

func GetSupplierByUserID(userID uuid.UUID) (*entity.Supplier, error) {
	return repository.GetSupplierByUserID(userID)
}

func UpdateSupplier(id uuid.UUID, req model.UpdateSupplierRequest) error {
	data := map[string]interface{}{}
	if req.StoreName != "" {
		data["store_name"] = req.StoreName
	}
	if req.Address != nil {
		data["address"] = req.Address
	}
	if req.ContactNumber != nil {
		data["contact_number"] = req.ContactNumber
	}
	if req.Category != nil {
		data["category"] = req.Category
	}
	if req.SourceType != nil {
		data["source_type"] = req.SourceType
	}
	if req.BusinessDesc != nil {
		data["business_desc"] = req.BusinessDesc
	}
	if req.NIBDocument != nil {
		data["nib_document"] = req.NIBDocument
	}
	if req.HalalDocument != nil {
		data["halal_document"] = req.HalalDocument
	}
	if req.OtherDocument != nil {
		data["other_document"] = req.OtherDocument
	}
	if req.AdminNotes != nil {
		data["admin_notes"] = req.AdminNotes
	}
	return repository.UpdateSupplier(id, data)
}

func VerifySupplier(id uuid.UUID, req model.VerifySupplierRequest) error {
	// Ambil data supplier
	supplier, err := repository.GetSupplierByID(id)
	if err != nil {
		return errors.New("supplier tidak ditemukan")
	}

	// Ambil data user untuk email
	user, err := repository.GetUserByID(supplier.UserID)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}

	// Update status verifikasi
	if err := repository.VerifySupplier(id, req.Status, req.Notes); err != nil {
		return err
	}

	// Kirim email notifikasi
	if req.Status == entity.VerificationApproved {
		go email.SendSupplierApprovedEmail(user.Email, user.Name, supplier.StoreName)
	} else if req.Status == entity.VerificationRejected {
		go email.SendSupplierRejectedEmail(user.Email, user.Name, supplier.StoreName, req.Notes)
	}

	return nil
}

func UpdateSupplierDocument(id uuid.UUID, docType string, path string) error {
	data := map[string]interface{}{}
	switch docType {
	case "nib":
		data["nib_document"] = path
	case "halal":
		data["halal_document"] = path
	case "other":
		data["other_document"] = path
	}
	return repository.UpdateSupplier(id, data)
}

func DeleteSupplier(id uuid.UUID) error {
	return repository.DeleteSupplier(id)
}