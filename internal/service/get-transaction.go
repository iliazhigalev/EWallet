package service

import (
	"context"

	"ewallet/internal/models"
)

func (e *EWallet) GetTransactionHistory(ctx context.Context, walletId string) ([]models.Transaction, error) {
	db := e.db
	history, err := e.transactionRepo.GetTransactionRecord(ctx, db, walletId)
	if err != nil {
		return nil, err
	}

	return history, nil
}
