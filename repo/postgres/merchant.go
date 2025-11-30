package pg

import (
	"context"

	models "github.com/SaeedSH02/payment_gateway/models"
	logger "github.com/SaeedSH02/payment_gateway/log"
)

func (p *postgres) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	exists, err := p.db.NewSelect().
		Model((*models.Merchant)(nil)).
		Where("email = ?", email).
		Exists(ctx)
	return exists, err
}
func (p *postgres) CreatMerchant(ctx context.Context, merchant *models.Merchant) error {
	if _, err := p.db.NewInsert().Model(merchant).Exec(ctx); err != nil {
		logger.Lg.Error("cant create merchant: ", "error: ", err)
		return err
	}
	logger.Lg.Info("user created: ", "user: ", merchant)
	return nil
}

func (p *postgres) GetMerchantByEmail(ctx context.Context, email string) (*models.Merchant, error) {
	merchant := new(models.Merchant)
	err := p.db.NewSelect().
		Model(merchant).
		Where("email = ?", email).
		Scan(ctx)
	if err != nil {
		logger.Lg.Error("cant get merchant by email: ", "error: ", err)
		return nil, err
	}
	return merchant, nil
}
