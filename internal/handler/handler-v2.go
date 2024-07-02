package handler

import (
	"context"
	"net/http"

	"ewallet/internal/models"
)

type eWalletService interface {
	CreateWallet(ctx context.Context) (models.Wallet, error)
	GetWallet(ctx context.Context, walletID string) (models.Wallet, error)
}

type Handler struct {
	eWalletService eWalletService
}

// GetTransactionHistory implements routes.handlerItem.
func (h *Handler) GetTransactionHistory(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// SendMoney implements routes.handlerItem.
func (h *Handler) SendMoney(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func New(
	eWalletService eWalletService,
) *Handler {
	return &Handler{
		eWalletService: eWalletService,
	}
}
