package rest

import (
	"sppg-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Global rate limiter untuk semua request
	r.Use(middleware.GlobalRateLimiter())

	r.Static("/uploads", "./uploads")

	api := r.Group("/api/v1")

	// Auth routes dengan rate limiter 
	authGroup := api.Group("")
	authGroup.Use(middleware.AuthRateLimiter())
	{
		AuthRoutes(authGroup)
		ForgotPasswordRoutes(authGroup)
	}

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
		ProductRoutes(protected)
		StockRoutes(protected)
		OrderRoutes(protected)
		TransactionRoutes(protected)
		PaymentRoutes(protected)
		
		// Upload dengan rate limiter khusus
		uploadGroup := protected.Group("")
		uploadGroup.Use(middleware.UploadRateLimiter())
		{
			UploadRoutes(uploadGroup)
		}
	}
}