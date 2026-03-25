package model

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	Token           string `json:"token"            binding:"required"`
	NewPassword     string `json:"new_password"     binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=6"`
}

type ForgotPasswordResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ForgotOK(message string) ForgotPasswordResponse {
	return ForgotPasswordResponse{Success: true, Message: message}
}

func ForgotFail(message string) ForgotPasswordResponse {
	return ForgotPasswordResponse{Success: false, Message: message}
}