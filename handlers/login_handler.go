package handlers

import (
	"github.com/SaeedSH02/payment_gateway/models"
	"github.com/SaeedSH02/payment_gateway/repo"
	"github.com/SaeedSH02/payment_gateway/service"
	"github.com/gin-gonic/gin"
)

func LoginHandler(db repo.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.Merchant_Login
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid request"})
			return
		}
		resp, err := service.LoginMerchant(c.Request.Context(),db, req)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{"token": resp.Token})
	}
}