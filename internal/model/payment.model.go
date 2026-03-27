package model

type CreatePaymentRequest struct {
	OrderID string `json:"order_id" binding:"required"`
}

type PaymentResponse struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data,omitempty"`
}

type MidtransNotification struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
	GrossAmount       string `json:"gross_amount"`
	SignatureKey      string `json:"signature_key"`
}

func PaymentOK(message string, data interface{}) PaymentResponse {
	return PaymentResponse{Success: true, Message: message, Data: data}
}

func PaymentFail(message string) PaymentResponse {
	return PaymentResponse{Success: false, Message: message}
}