package pg

import (
	"context"

	model "github.com/SaeedSH02/payment_gateway/Models"
	logger "github.com/SaeedSH02/payment_gateway/log"
	"go.uber.org/zap"
)

func (p *postgres) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	exists, err := p.db.NewSelect().
		Model((*model.Merchant)(nil)).
		Where("email = ?", email).
		Exists(ctx)
	return exists, err
}
func (p *postgres) CreatMerchant(ctx context.Context, merchant *model.Merchant) error {
	if _, err := p.db.NewInsert().Model(merchant).Exec(ctx); err != nil {
		logger.LogError(ctx, "cant create merchant", err)
		return err
	}
	logger.FromContext(ctx).Info("user created: ", zap.Any("user: ", merchant))
	return nil
}
