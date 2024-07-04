package handler

import (
	"context"

	"ewallet/internal/dto"
	"ewallet/internal/models"
)

type eWalletService interface {
	CreateWallet(ctx context.Context) (models.Wallet, error)
	GetWallet(ctx context.Context, walletID string) (models.Wallet, error)
	SendMoney(ctx context.Context, walletID string, request dto.SendRequest) error
	GetTransactionHistory(ctx context.Context, walletId string) ([]models.Transaction, error)
}

type Handler struct {
	eWalletService eWalletService
}

func New(
	eWalletService eWalletService,
) *Handler {
	return &Handler{
		eWalletService: eWalletService,
	}
}
