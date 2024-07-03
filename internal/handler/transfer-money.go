package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"ewallet/internal/dto"

	"github.com/gorilla/mux"
)

func (h Handler) SendMoney(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)
	walletId := vars["walletId"]

	var req dto.SendRequest // содержит id куда надо перевести деньги

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.eWalletService.SendMoney(ctx, walletId, req)
	if err != nil {
		if errors.Is(err, dto.ErrWalletNotFound) {
			http.Error(w, "From wallet not found", http.StatusNotFound)
			return
		} else if errors.Is(err, dto.ErrInsufficientFunds) {
			http.Error(w, "Insufficient funds for transfer to another wallet", http.StatusBadRequest)
			return
		} else if errors.Is(err, dto.ErrTargetWalletNotFound) {
			http.Error(w, "The target wallet was not found", http.StatusBadRequest)
			return
		} else {
			http.Error(w, "the server is not responding", http.StatusInternalServerError)
			return
		}

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}
