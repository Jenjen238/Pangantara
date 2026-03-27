package app

import (
	"sppg-backend/internal/controller/rest"
	ginpkg "sppg-backend/pkg/gin"
	"sppg-backend/pkg/payment"
	"sppg-backend/pkg/postgres"
)

func Run() {
	postgres.Connect()
	postgres.Migrate()
	payment.InitMidtrans()

	ginpkg.Init()
	rest.RegisterRoutes(ginpkg.Router)
	ginpkg.Run()
}