package model

import (
	"sppg-backend/internal/entity"

	"github.com/google/uuid"
)

type CreateSupplierRequest struct {
	UserID        uuid.UUID `json:"user_id"        binding:"required"`
	StoreName     string    `json:"store_name"     binding:"required"`
	Address       *string   `json:"address"`
	ContactNumber *string   `json:"contact_number"`
	Category      *string   `json:"category"`
	SourceType    *string   `json:"source_type"`
	BusinessDesc  *string   `json:"business_desc"`
	NIBDocument   *string   `json:"nib_document"`
	HalalDocument *string   `json:"halal_document"`
	OtherDocument *string   `json:"other_document"`
	AdminNotes    *string   `json:"admin_notes"`
}

type UpdateSupplierRequest struct {
	StoreName     string  `json:"store_name"`
	Address       *string `json:"address"`
	ContactNumber *string `json:"contact_number"`
	Category      *string `json:"category"`
	SourceType    *string `json:"source_type"`
	BusinessDesc  *string `json:"business_desc"`
	NIBDocument   *string `json:"nib_document"`
	HalalDocument *string `json:"halal_document"`
	OtherDocument *string `json:"other_document"`
	AdminNotes    *string `json:"admin_notes"`
}

type VerifySupplierRequest struct {
	Status entity.VerificationStatus `json:"status" binding:"required,oneof=approved rejected"`
	Notes  *string                   `json:"admin_notes"`
}

type SupplierResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SupplierOK(message string, data interface{}) SupplierResponse {
	return SupplierResponse{Success: true, Message: message, Data: data}
}

func SupplierFail(message string) SupplierResponse {
	return SupplierResponse{Success: false, Message: message}
}