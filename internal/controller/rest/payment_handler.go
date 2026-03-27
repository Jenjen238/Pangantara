package rest

import (
	"net/http"
	"sppg-backend/internal/middleware"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.RouterGroup) {
	payment := r.Group("/payment")
	{
		// Protected - SPPG buat payment
		payment.POST("/create", middleware.RoleMiddleware("sppg", "admin"), createPayment)
	}
}

func WebhookRoutes(r *gin.RouterGroup) {
	webhook := r.Group("/webhook")
	{
		// Public - Midtrans callback
		webhook.POST("/midtrans", midtransWebhook)
	}
}

func createPayment(c *gin.Context) {
	var req model.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.PaymentFail(err.Error()))
		return
	}

	resp, err := usecase.CreatePayment(req.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.PaymentFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.PaymentOK("Payment berhasil dibuat", gin.H{
		"token":        resp.Token,
		"redirect_url": resp.RedirectURL,
	}))
}

func midtransWebhook(c *gin.Context) {
	var notif model.MidtransNotification
	if err := c.ShouldBindJSON(&notif); err != nil {
		c.JSON(http.StatusBadRequest, model.PaymentFail(err.Error()))
		return
	}

	if err := usecase.HandleMidtransNotification(notif); err != nil {
		c.JSON(http.StatusBadRequest, model.PaymentFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.PaymentOK("Notifikasi berhasil diproses", nil))
}