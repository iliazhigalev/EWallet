package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) GetTransactionHistory(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)
	walletId := vars["walletId"]
	_, err := h.eWalletService.GetWallet(ctx, walletId)
	if err != nil {
		http.Error(w, fmt.Sprintf("create wallet: %s", err), http.StatusNotFound)
		return
	}
	history, err := h.eWalletService.GetTransactionHistory(ctx, walletId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(history)
}
