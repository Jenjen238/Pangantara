package gin

import (
	"log"
	"sppg-backend/config"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Init() {
	if config.AppConfig.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	Router = gin.Default()

	allowedOrigins := []string{
    "http://localhost:8080",           
    "https://be-internship.bccdev.id/hanif", 
	"https://pangantara.vercel.app",        
}

Router.Use(func(c *gin.Context) {
    origin := c.Request.Header.Get("Origin")
    for _, allowed := range allowedOrigins {
        if origin == allowed {
            c.Header("Access-Control-Allow-Origin", origin)
            break
        }
    }
    c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
    c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")
    c.Header("Access-Control-Allow-Credentials", "true")
    c.Header("Access-Control-Expose-Headers", "Content-Length")
    if c.Request.Method == "OPTIONS" {
        c.AbortWithStatus(204)
        return
    }
    c.Next()
})

	Router.MaxMultipartMemory = 10 << 20 // 10 MB

	Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong", "status": "ok"})
	})
}

func Run() {
	port := config.AppConfig.AppPort
	log.Printf("Server berjalan di http://localhost:%s", port)
	if err := Router.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}