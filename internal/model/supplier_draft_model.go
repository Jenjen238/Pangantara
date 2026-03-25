package model

import "github.com/google/uuid"

type SaveDraftRequest struct {
	UserID        uuid.UUID `json:"user_id"        binding:"required"`
	StoreName     *string   `json:"store_name"`
	Address       *string   `json:"address"`
	ContactNumber *string   `json:"contact_number"`
	Category      *string   `json:"category"`
	SourceType    *string   `json:"source_type"`
	BusinessDesc  *string   `json:"business_desc"`
	CurrentStep   int       `json:"current_step"`
}

type SubmitDraftRequest struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`
}

type DraftResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func DraftOK(message string, data interface{}) DraftResponse {
	return DraftResponse{Success: true, Message: message, Data: data}
}

func DraftFail(message string) DraftResponse {
	return DraftResponse{Success: false, Message: message}
}