package handler

import (
	"context"
	"net/http"

	"ewallet/internal/dto"
	"ewallet/internal/models"
)

type eWalletService interface {
	CreateWallet(ctx context.Context) (models.Wallet, error)
	GetWallet(ctx context.Context, walletID string) (models.Wallet, error)
	SendMoney(ctx context.Context, walletID string, request dto.SendRequest) error
}

type Handler struct {
	eWalletService eWalletService
}

func (h Handler) GetTransactionHistory(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func New(
	eWalletService eWalletService,
) *Handler {
	return &Handler{
		eWalletService: eWalletService,
	}
}
