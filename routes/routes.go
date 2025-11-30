package routes

import (
	"time"

	"github.com/SaeedSH02/payment_gateway/handlers"
	logger "github.com/SaeedSH02/payment_gateway/log"
	"github.com/SaeedSH02/payment_gateway/middleware"
	"github.com/SaeedSH02/payment_gateway/repo"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func SetupRoutes(db repo.DB) *gin.Engine {
	r := gin.Default()

	r.Use(sloggin.New(logger.Lg))
	r.Use(gin.Recovery())
	r.Use(timeout.New(
		timeout.WithTimeout(15 * time.Second),
		timeout.WithResponse(func(c *gin.Context) {
			c.JSON(503, gin.H{"error": "request timed out"})
			logger.Lg.Warn("request timed out", "path", c.Request.URL.Path)
		}),
	))

	r.POST("/register", handlers.RegisterHandler(db))
	r.POST("/login", handlers.LoginHandler(db))
	r.GET("/test", func(ctx *gin.Context) {
		time.Sleep(10 * time.Second)
	})

	//Merchant Routes
	protected := r.Group("/merchant")
	protected.Use(middleware.JWTMiddleware())
	{
		protected.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"test": "Successful"})
		})
	}

	return r
}
