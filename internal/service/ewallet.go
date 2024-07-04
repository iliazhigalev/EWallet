package service

import (
	"context"

	"ewallet/internal/models"
)

type walletRepository interface {
	Create(ctx context.Context, db DB, wallet models.Wallet) error
	GetByWalletID(ctx context.Context, db DB, walletID string) (models.Wallet, error)
	UpdateBalance(ctx context.Context, db DB, walletID string, balance float64) error
}

type transactionRepository interface {
	Create(ctx context.Context, db DB, wallet models.Transaction) error
	GetTransactionRecord(ctx context.Context, db DB, walletID string) ([]models.Transaction, error)
}

type EWallet struct {
	db              transactionalDB
	walletRepo      walletRepository
	transactionRepo transactionRepository
}

func NewEWallet(db transactionalDB, walletRepo walletRepository, transactionRepo transactionRepository) *EWallet {
	return &EWallet{
		db:              db,
		walletRepo:      walletRepo,
		transactionRepo: transactionRepo,
	}
}
