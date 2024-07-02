package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) GetWallet(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)
	walletId := vars["walletId"]
	wallet, err := h.eWalletService.GetWallet(ctx, walletId)
	if err != nil {
		http.Error(w, fmt.Sprintf("create wallet: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wallet)
}
