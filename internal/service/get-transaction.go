package service

import (
	"context"

	"ewallet/internal/models"
)

func (e EWallet) GetTransactionHistory(ctx context.Context, walletId string) ([]models.Transaction, error) {
	tx, err := e.db.Begin()
	if err != nil {
		return nil, err
	}
	history, err := e.transactionRepo.GetTransactionRecord(ctx, tx, walletId)
	if err == nil {
		return nil, err
	}

	return history, err
}
