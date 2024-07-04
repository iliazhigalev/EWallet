package service

import (
	"context"
	"database/sql"
)

type DB interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) // для создания
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row        // для получения
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

type transactionalDB interface {
	DB
	Begin() (*sql.Tx, error)
}
