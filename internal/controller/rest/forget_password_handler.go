package rest

import (
	"net/http"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

func ForgotPasswordRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/forgot-password", forgotPassword)
		auth.POST("/reset-password", resetPassword)
	}
}

func forgotPassword(c *gin.Context) {
	var req model.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ForgotFail(err.Error()))
		return
	}
	if err := usecase.ForgotPassword(req); err != nil {
		c.JSON(http.StatusInternalServerError, model.ForgotFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.ForgotOK("Jika email terdaftar, link reset password akan dikirim"))
}

func resetPassword(c *gin.Context) {
	var req model.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ForgotFail(err.Error()))
		return
	}
	if err := usecase.ResetPassword(req); err != nil {
		c.JSON(http.StatusBadRequest, model.ForgotFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.ForgotOK("Kata sandi berhasil direset"))
}