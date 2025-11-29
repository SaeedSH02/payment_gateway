package routes

import (
	"github.com/SaeedSH02/payment_gateway/handlers"
	"github.com/SaeedSH02/payment_gateway/repo"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(db repo.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/register", handlers.RegisterHandler(db))


	return r
}