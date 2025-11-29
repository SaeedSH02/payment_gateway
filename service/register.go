package service

import (
	"context"
	"errors"

	model "github.com/SaeedSH02/payment_gateway/Models"
	logger "github.com/SaeedSH02/payment_gateway/log"
	"github.com/SaeedSH02/payment_gateway/repo"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	return string(bytes), err
}



func RegisterMerchant(ctx context.Context, db repo.DB, input *model.Merchant_Input) error {
	exist, err := db.ExistsByEmail(ctx, input.Email)
	if err != nil {
		logger.LogError(ctx, "error while cheking Email", err)
		return  err
	}
	if exist == true {
		return errors.New("email already existed")
	}

	hash, err := hashPassword(input.Password)
	if err != nil {
		logger.LogError(ctx, "error while hashing password..", err)
		return err
	}
	merchant := model.Merchant{
		Name: input.Name,
		Email: input.Email,
		PasswordHash: hash,
	}

	err = db.CreatMerchant(ctx, &merchant)
	return err

}


