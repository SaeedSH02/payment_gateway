package service

import (
	"context"
	"errors"
	"os"
	"time"

	logger "github.com/SaeedSH02/payment_gateway/log"
	models "github.com/SaeedSH02/payment_gateway/models"
	"github.com/SaeedSH02/payment_gateway/repo"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)



func LoginMerchant(ctx context.Context, db repo.DB, req models.Merchant_Login) (models.LoginResponse, error) {
	//Get email and password from request body
	merchant, err := db.GetMerchantByEmail(ctx, req.Email)
	if err != nil {
		logger.Lg.Error("error getting merchant by email: ", "error: ", err, "email: ", req.Email)
		return models.LoginResponse{}, errors.New("invalid email or password")
	}
	//create JWT token

	if bcrypt.CompareHashAndPassword([]byte(merchant.PasswordHash), []byte(req.Password)) != nil {
		logger.Lg.Warn("invalid password for email: ", "email: ", req.Email)
		return models.LoginResponse{}, errors.New("invalid email or password")
	}
	token, err := generateToken(merchant.Email)
	if err != nil {
		logger.Lg.Error("error generating token: ", "error: ", err)
		return models.LoginResponse{}, errors.New("could not generate token")
	}
	logger.Lg.Info("merchant logged in successfully: ", "email: ", req.Email)
	return models.LoginResponse{Token: token}, nil
}



func generateToken(email string) (string, error) {
	expirationTime := time.Now().Add(48 * time.Hour)
	claims := &models.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}