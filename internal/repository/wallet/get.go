package wallet

import (
	"context"
	"ewallet/internal/dto"
	"ewallet/internal/models"
	"ewallet/internal/service"
)

func (w *Wallet) GetByWalletID(ctx context.Context, db service.DB, walletID string) (models.Wallet, error) {
	var wallet models.Wallet
	query := `SELECT id, balance FROM wallets WHERE id = $1`
	row := db.QueryRowContext(ctx, query, walletID)

	err := row.Scan(&wallet.ID, &wallet.Balance)
	if err != nil {
		return wallet, dto.ErrNotFound
	}

	return wallet, nil
}
