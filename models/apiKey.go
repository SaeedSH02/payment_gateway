package model

import (
	"time"
	"github.com/uptrace/bun"
)




type ApiKey struct {
	bun.BaseModel `bun:"table:api_keys,alias:ak"`

	ID         int64     `bun:"id,pk,autoincrement"`
	MerchantID int64     `bun:"merchant_id, notnull"`
	ApiKey     string    `bun:"api_key,unique,notnull"`
	SecretKey  string    `bun:"secret_key, notnull"`
	IsActive   bool      `bun:"is_active,default:true"`
	CreatedAt  time.Time `bun:"created_at,nullzero,default:now()"`

	Merchant *Merchant `bun:"rel:belongs-to,join:merchant_id=id"`
}

