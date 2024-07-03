package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"ewallet/internal/service"

	"github.com/gorilla/mux"
)

type SendRequest struct {
	To     string  `json:"to"`     // id куда нужна перевести деньги
	Amount float64 `json:"amount"` // сумма перевода
}

func (h *Handler) SendMoney(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()
	vars := mux.Vars(r)
	walletId := vars["walletId"]

	var req SendRequest // содержит id куда надо перевести деньги

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return err
	}

	err := h.eWalletService.SendMoney(ctx, walletId, req)
	if err != nil {
		if errors.Is(err, service.ErrWalletNotFound) {
			http.Error(w, "From wallet not found", http.StatusNotFound)
			return err
		} else if errors.Is(err, service.ErrInsufficientFunds) {
			http.Error(w, "Insufficient funds for transfer to another wallet", http.StatusBadRequest)
			return err
		} else if errors.Is(err, service.ErrTargetWalletNotFound) {
			http.Error(w, "The target wallet was not found", http.StatusBadRequest)
			return err
		} else {
			http.Error(w, "the server is not responding", http.StatusInternalServerError)
			return err
		}

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return nil
}
