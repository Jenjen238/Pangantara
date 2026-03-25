package rest

import (
	"sppg-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.Static("/uploads", "./uploads")

	api := r.Group("/api/v1")

	AuthRoutes(api)
	ForgotPasswordRoutes(api)

	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		DashboardRoutes(protected)
		UserRoutes(protected)
		SPPGRoutes(protected)
		SupplierRoutes(protected)
		UploadRoutes(protected)
		ProductRoutes(protected)
		StockRoutes(protected)
		OrderRoutes(protected)
		TransactionRoutes(protected)
		SupplierDraftRoutes(protected)
	}
}