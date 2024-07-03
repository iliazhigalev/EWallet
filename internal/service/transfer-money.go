package service

import (
	"context"
	"ewallet/internal/dto"
	"ewallet/internal/models"
	"ewallet/pkg/utils"
	"time"
)

func (e EWallet) SendMoney(ctx context.Context, walletID string, request dto.SendRequest) error {
	tx, err := e.db.Begin()
	if err != nil {
		return err
	}
	fromWallet, err := e.walletRepo.GetByWalletID(ctx, tx, walletID)
	if err == dto.ErrNotFound {
		return dto.ErrWalletNotFound
	}
	toWallet, err := e.walletRepo.GetByWalletID(ctx, tx, request.To)
	if err == dto.ErrNotFound {
		return dto.ErrTargetWalletNotFound
	}
	if fromWallet.Balance < request.Balance {
		return dto.ErrInsufficientFunds
	} else {
		fromBalance := fromWallet.Balance - request.Balance
		err = e.walletRepo.UpdateBalance(ctx, tx, fromWallet.ID, fromBalance)
		if err != nil {
			tx.Rollback()
			return err
		}
		toBalance := toWallet.Balance + request.Balance
		err := e.walletRepo.UpdateBalance(ctx, tx, toWallet.ID, toBalance)
		if err != nil {
			tx.Rollback()
			return err
		}
		err = e.transactionRepo.Create(ctx, tx, models.Transaction{
			ID:         utils.GenerateID(),
			Time:       time.Now(),
			FromWallet: fromWallet.ID,
			ToWallet:   toWallet.ID,
			Amount:     initBalance,
		})
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}
