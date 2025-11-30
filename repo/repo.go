package repo

import (
	"context"

	models "github.com/SaeedSH02/payment_gateway/models"
)

type DB interface {
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	CreatMerchant(ctx context.Context, merchant *models.Merchant) error
	GetMerchantByEmail(ctx context.Context, email string) (*models.Merchant, error)
}