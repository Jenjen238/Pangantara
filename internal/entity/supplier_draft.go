package entity

import (
	"github.com/google/uuid"
	"time"
)

type DraftStatus string

const (
	DraftStatusDraft     DraftStatus = "draft"
	DraftStatusSubmitted DraftStatus = "submitted"
)

type SupplierDraft struct {
	ID            uuid.UUID   `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID        uuid.UUID   `gorm:"type:uuid;not null;uniqueIndex"                  json:"user_id"`
	StoreName     *string     `gorm:"type:varchar(150)"                               json:"store_name,omitempty"`
	Address       *string     `gorm:"type:varchar"                                    json:"address,omitempty"`
	ContactNumber *string     `gorm:"type:varchar(20)"                                json:"contact_number,omitempty"`
	Category      *string     `gorm:"type:varchar(50)"                                json:"category,omitempty"`
	SourceType    *string     `gorm:"type:varchar(50)"                                json:"source_type,omitempty"`
	BusinessDesc  *string     `gorm:"type:text"                                       json:"business_desc,omitempty"`
	NIBDocument   *string     `gorm:"type:varchar(255)"                               json:"nib_document,omitempty"`
	HalalDocument *string     `gorm:"type:varchar(255)"                               json:"halal_document,omitempty"`
	OtherDocument *string     `gorm:"type:varchar(255)"                               json:"other_document,omitempty"`
	CurrentStep   int         `gorm:"not null;default:1"                              json:"current_step"`
	Status        DraftStatus `gorm:"type:varchar(20);not null;default:'draft'"       json:"status"`
	CreatedAt     time.Time   `gorm:"autoCreateTime"                                  json:"created_at"`
	UpdatedAt     time.Time   `gorm:"autoUpdateTime"                                  json:"updated_at"`
}

func (SupplierDraft) TableName() string { return "supplier_drafts" }