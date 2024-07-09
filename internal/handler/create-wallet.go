package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateWallet(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()
	wallet, err := h.eWalletService.CreateWallet(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("create wallet: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wallet)
}
