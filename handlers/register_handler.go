package handlers

import (
	model "github.com/SaeedSH02/payment_gateway/Models"
	"github.com/SaeedSH02/payment_gateway/repo"
	"github.com/SaeedSH02/payment_gateway/service"
	"github.com/gin-gonic/gin"
)



func RegisterHandler(db repo.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var merchant model.Merchant_Input

		if err := c.ShouldBindJSON(&merchant); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err := service.RegisterMerchant(c, db, &merchant)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return 
		}

		c.JSON(200, gin.H{"message": "registered successfully"})
	}
}