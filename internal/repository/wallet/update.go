package wallet

import (
	"context"

	"ewallet/internal/service"
)

func (w *Wallet) UpdateBalance(ctx context.Context, db service.DB, walletID string, balance float64) error {
	query := `IPDATE wallets SET balance = $1 WHERE id = $2`
	_, err := db.ExecContext(ctx, query, balance, walletID)

	return err
}
