package service

import (
	"context"
	"fmt"
	"time"

	"ewallet/internal/models"
	"ewallet/pkg/utils"
)

const (
	initBalance = float64(100)
	adminID     = "admin1"
)

func (e EWallet) CreateWallet(ctx context.Context) (m models.Wallet, er error) {
	tx, err := e.db.Begin()
	if err != nil {
		return models.Wallet{}, err
	}
	defer func(e error) {
		if e != nil {
			tx.Rollback()
		}
		tx.Commit()
	}(er)

	wallet := models.Wallet{
		ID:      utils.GenerateID(),
		Balance: initBalance,
	}
	err = e.walletRepo.Create(ctx, tx, wallet)
	if err != nil {
		tx.Rollback()
		return models.Wallet{}, fmt.Errorf("create wallet: %w", err)
	}

	err = e.transactionRepo.Create(ctx, tx, models.Transaction{
		ID:         utils.GenerateID(),
		Time:       time.Now(),
		FromWallet: adminID,
		ToWallet:   wallet.ID,
		Amount:     initBalance,
	})
	if err != nil {
		tx.Rollback()
		return models.Wallet{}, fmt.Errorf("create transaction: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return models.Wallet{}, fmt.Errorf("commit store transaction: %w", err)
	}

	return wallet, nil
}
