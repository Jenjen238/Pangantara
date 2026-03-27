package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

var (
	// Global rate limit — 100 request per menit
	globalStore = memory.NewStore()
	globalRate  = limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  100,
	}

	// Auth rate limit — 10 request per menit (anti brute force)
	authStore = memory.NewStore()
	authRate  = limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  10,
	}

	// Upload rate limit — 20 request per menit
	uploadStore = memory.NewStore()
	uploadRate  = limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  20,
	}
)

func RateLimiter(store limiter.Store, rate limiter.Rate) gin.HandlerFunc {
	instance := limiter.New(store, rate)
	return func(c *gin.Context) {
		context, err := instance.Get(c.Request.Context(), c.ClientIP())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Rate limiter error",
			})
			c.Abort()
			return
		}

		// Set header info rate limit
		c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", context.Limit))
		c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", context.Remaining))
		c.Header("X-RateLimit-Reset", fmt.Sprintf("%d", context.Reset))

		if context.Reached {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"message": "Terlalu banyak request, coba lagi nanti",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func GlobalRateLimiter() gin.HandlerFunc {
	return RateLimiter(globalStore, globalRate)
}

func AuthRateLimiter() gin.HandlerFunc {
	return RateLimiter(authStore, authRate)
}

func UploadRateLimiter() gin.HandlerFunc {
	return RateLimiter(uploadStore, uploadRate)
}