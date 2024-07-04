package transaction

import (
	"context"
	"ewallet/internal/models"
	"ewallet/internal/service"
)

func (w *Transaction) GetTransactionRecord(ctx context.Context, db service.DB, walletId string) ([]models.Transaction, error) {
	query := "SELECT id, time, from_wallet, to_wallet, amount FROM transactions WHERE from_wallet= $1 OR to_wallet=$1"
	history, err := db.QueryContext(ctx, query, walletId)
	if err != nil {
		return nil, err
	}
	var transactions []models.Transaction
	for history.Next() {
		var transaction models.Transaction
		if err := history.Scan(&transaction.ID, &transaction.Time, &transaction.FromWallet, &transaction.ToWallet, &transaction.Amount); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
