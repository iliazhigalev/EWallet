package transaction

import (
	"context"

	"ewallet/internal/models"
	"ewallet/internal/service"
	"ewallet/pkg/utils"
)

func (w *Transaction) Create(ctx context.Context, db service.DB, transaction models.Transaction) error {
	transactionQuery := `INSERT INTO transactions (id, time, from_wallet, to_wallet, amount) VALUES ($1, $2, $3, $4, $5)`
	transactionID := utils.GenerateID()
	_, err := db.ExecContext(
		ctx,
		transactionQuery,
		transactionID,
		transaction.Time,
		transaction.FromWallet,
		transaction.ToWallet,
		transaction.Amount,
	)

	return err
}
