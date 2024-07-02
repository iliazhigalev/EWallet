package wallet

import (
	"context"
	"database/sql"
	"ewallet/internal/models"
	"ewallet/internal/service"
)

func (w *Wallet) GetByWalletID(ctx context.Context, db service.DB, walletID string) (models.Wallet, error) {
	var wallet models.Wallet
	query := `SELECT id, balance FROM wallets WHERE id = $1`
	row := db.QueryRowContext(ctx, query, walletID)

	err := row.Scan(&wallet.ID, &wallet.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return wallet, nil
		}
		return wallet, err
	}

	return wallet, nil
}
