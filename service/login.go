package service

import (
	"os"
	"time"

	logger "github.com/SaeedSH02/payment_gateway/log"
	models "github.com/SaeedSH02/payment_gateway/models"
	"github.com/SaeedSH02/payment_gateway/repo"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)



func LoginMerchant(repo repo.DB, C *gin.Context) {
	//Get email and password from request body
	var merchant models.Merchant_Login
	if err := C.ShouldBindJSON(&merchant); err != nil {
		C.JSON(400, gin.H{"error": "Invalid request"})
		logger.Lg.Warn("invalid login request: ", "error: ", err, "email: ", merchant.Email)
		return
	}
	//lookup user in database
	merchantFromDb, err := repo.GetMerchantByEmail(C.Request.Context(), merchant.Email)
	if err != nil {
		C.JSON(401, gin.H{"error": "Invalid email or password"})
		logger.Lg.Warn("invalid email: ", "email: ", merchant.Email)
		return
	}
	//compare password with hashed password
	err = bcrypt.CompareHashAndPassword([]byte(merchantFromDb.PasswordHash), []byte(merchant.Password))
	if err != nil {
		C.JSON(401, gin.H{"error": "Invalid email or password"})
		logger.Lg.Warn("invalid password for email: ", "email: ", merchant.Email)
		return
	}
	//create JWT token
	expirationTime := time.Now().Add(48 * time.Hour)
	claims := &models.Claims{
		Email: merchantFromDb.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		C.JSON(500, gin.H{"error": "internal Error Try later"})
		logger.Lg.Error("cant sign token: ", "error: ", err, "email: ", merchant.Email)
		return
	}

	//return token in response
	logger.Lg.Info("merchant logged in: ", "email: ", merchant.Email)
	C.JSON(200, gin.H{"token": tokenString})
}
