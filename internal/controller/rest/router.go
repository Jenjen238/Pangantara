package rest

import (
	"sppg-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.Use(middleware.GlobalRateLimiter())
	r.Static("/uploads", "./uploads")

	api := r.Group("/api/v1")

	// Public routes
	authGroup := api.Group("")
	authGroup.Use(middleware.AuthRateLimiter())
	{
		AuthRoutes(authGroup)
		ForgotPasswordRoutes(authGroup)
	}

	// Webhook
	WebhookRoutes(api)

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		DashboardRoutes(protected)
		UserRoutes(protected)
		SPPGRoutes(protected)
		SupplierRoutes(protected)
		SupplierDraftRoutes(protected)
		OrderRoutes(protected)
		TransactionRoutes(protected)
		PaymentRoutes(protected)

		// Endpoint yang butuh supplier verified
		supplierVerified := protected.Group("")
		supplierVerified.Use(middleware.SupplierVerifiedMiddleware())
		{
			ProductRoutes(supplierVerified)
			StockRoutes(supplierVerified)

			uploadGroup := supplierVerified.Group("")
			uploadGroup.Use(middleware.UploadRateLimiter())
			{
				UploadRoutes(uploadGroup)
			}
		}
	}
}