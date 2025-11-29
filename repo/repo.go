package repo

import (
	"context"

	model "github.com/SaeedSH02/payment_gateway/Models"
)

type DB interface {
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	CreatMerchant(ctx context.Context, merchant *model.Merchant) error
}