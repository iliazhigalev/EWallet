package handler

import (
	"context"

	"ewallet/internal/models"
)

type eWalletService interface {
	CreateWallet(ctx context.Context) (models.Wallet, error)
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
