package service

import (
	"context"

	"ewallet/internal/models"
)

func (e EWallet) GetWallet(ctx context.Context, walletID string) (models.Wallet, error) {
	tx, err := e.db.Begin()
	if err != nil {
		return models.Wallet{}, err
	}
	wallet, err := e.walletRepo.GetByWalletID(ctx, tx, walletID)
	if err == nil {
		return wallet, err
	}

	return wallet, err
}
