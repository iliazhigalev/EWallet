package wallet

import (
	"context"

	"ewallet/internal/models"
	"ewallet/internal/service"
)

func (w *Wallet) Create(ctx context.Context, db service.DB, wallet models.Wallet) error {
	query := `INSERT INTO wallets (id, balance) VALUES ($1, $2)`
	_, err := db.ExecContext(ctx, query, wallet.ID, wallet.Balance)

	return err
}
