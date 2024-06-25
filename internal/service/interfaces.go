package service

import (
	"context"
	"database/sql"
)

type DB interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type transactionalDB interface {
	DB
	Begin() (*sql.Tx, error)
}
