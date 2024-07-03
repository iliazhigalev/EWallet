package service

import (
	"context"
	"errors"
	"ewallet/internal/handler"
	"ewallet/internal/models"
	"ewallet/internal/repository/wallet"
	"ewallet/pkg/utils"
	"time"
)

var (
	ErrWalletNotFound       = errors.New("outgoing wallet not found")
	ErrInsufficientFunds    = errors.New("insufficient funds")
	ErrTargetWalletNotFound = errors.New("target wallet not found")
)

type SendRequest struct {
	To     string  `json:"to"`     // id куда нужна перевести деньги
	Amount float64 `json:"amount"` // сумма перевода
}

func (e EWallet) SendMoney(ctx context.Context, walletID string, request handler.SendRequest) error {
	tx, err := e.db.Begin()
	if err != nil {
		return err
	}
	fromWallet, err := e.walletRepo.GetByWalletID(ctx, tx, walletID)
	if err == wallet.ErrNotFound {
		return ErrWalletNotFound
	}
	toWallet, err := e.walletRepo.GetByWalletID(ctx, tx, request.To)
	if err == wallet.ErrNotFound {
		return ErrTargetWalletNotFound
	}
	if fromWallet.Balance < request.Amount {
		return ErrInsufficientFunds
	} else {
		fromBalance := fromWallet.Balance - request.Amount
		err = e.walletRepo.UpdateBalance(ctx, tx, fromWallet.ID, fromBalance)
		if err != nil {
			tx.Rollback()
			return err
		}
		toBalance := toWallet.Balance + request.Amount
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
