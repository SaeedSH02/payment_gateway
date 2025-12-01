package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/uptrace/bun"
)

type Merchant struct {
	bun.BaseModel `bun:"table:merchants,alias:m"`

	ID           int64     `bun:"id,pk,autoincrement"`
	Name         string    `bun:"name,notnull"`
	Email        string    `bun:"email,unique,notnull"`
	PasswordHash string    `bun:"password_hash,notnull"`
	WebhookURL   string    `bun:"webhook_url"`
	CreatedAt    time.Time `bun:"created_at,nullzero,default:now()"`

	ApiKeys []*ApiKey `bun:"rel:has-many,join:id=merchant_id"`
}

type Merchant_Input struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type Merchant_Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}
type LoginResponse struct {
    Token string `json:"token"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
