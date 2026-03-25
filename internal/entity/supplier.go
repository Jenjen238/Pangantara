package entity

import "github.com/google/uuid"

type VerificationStatus string

const (
	VerificationPending  VerificationStatus = "pending"
	VerificationApproved VerificationStatus = "approved"
	VerificationRejected VerificationStatus = "rejected"
)

type Supplier struct {
	Base
	SupplierID         uuid.UUID          `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"supplier_id"`
	UserID             uuid.UUID          `gorm:"type:uuid;not null"                              json:"user_id"`
	StoreName          string             `gorm:"type:varchar(150);not null"                      json:"store_name"`
	Address            *string            `gorm:"type:varchar"                                    json:"address,omitempty"`
	ContactNumber      *string            `gorm:"type:varchar(20)"                                json:"contact_number,omitempty"`
	Category           *string            `gorm:"type:varchar(50)"                                json:"category,omitempty"`
	SourceType         *string            `gorm:"type:varchar(50)"                                json:"source_type,omitempty"`
	BusinessDesc       *string            `gorm:"type:text"                                       json:"business_desc,omitempty"`
	NIBDocument        *string            `gorm:"type:varchar(255)"                               json:"nib_document,omitempty"`
	HalalDocument      *string            `gorm:"type:varchar(255)"                               json:"halal_document,omitempty"`
	OtherDocument      *string            `gorm:"type:varchar(255)"                               json:"other_document,omitempty"`
	VerificationStatus VerificationStatus `gorm:"type:varchar(20);not null;default:'pending'"     json:"verification_status"`
	AdminNotes         *string            `gorm:"type:text"                                       json:"admin_notes,omitempty"`
}

func (Supplier) TableName() string {
	return "suppliers"
}