package middleware

import (
	"net/http"
	"sppg-backend/internal/model"
	"sppg-backend/internal/repository"
	"sppg-backend/internal/entity"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SupplierVerifiedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")

		// Hanya cek kalau role supplier
		if role != "supplier" {
			c.Next()
			return
		}

		userIDStr := c.GetString("user_id")
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.Response{
				Success: false,
				Message: "Invalid user ID",
			})
			c.Abort()
			return
		}

		// Cek status verifikasi supplier
		supplier, err := repository.GetSupplierByUserID(userID)
		if err != nil {
			c.JSON(http.StatusForbidden, model.Response{
				Success: false,
				Message: "Supplier profile not found, please complete registration first",
			})
			c.Abort()
			return
		}

		if supplier.VerificationStatus == entity.VerificationPending {
			c.JSON(http.StatusForbidden, model.Response{
				Success: false,
				Message: "Your account is pending verification by admin, please wait",
			})
			c.Abort()
			return
		}

		if supplier.VerificationStatus == entity.VerificationRejected {
			c.JSON(http.StatusForbidden, model.Response{
				Success: false,
				Message: "Your account has been rejected, please contact admin",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}