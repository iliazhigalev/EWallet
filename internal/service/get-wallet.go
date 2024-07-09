package service

import (
	"context"

	"ewallet/internal/models"
)

func (e *EWallet) GetWallet(ctx context.Context, walletID string) (models.Wallet, error) {
	db := e.db

	wallet, err := e.walletRepo.GetByWalletID(ctx, db, walletID)
	if err == nil {
		return wallet, err
	}

	return wallet, err
}
