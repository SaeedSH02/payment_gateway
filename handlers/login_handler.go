package handlers

import (
	"github.com/SaeedSH02/payment_gateway/repo"
	"github.com/SaeedSH02/payment_gateway/service"
	"github.com/gin-gonic/gin"
)

func LoginHandler(db repo.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		service.LoginMerchant(db, c)
	}
}