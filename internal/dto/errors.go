package dto

import "errors"

var (
	ErrWalletNotFound       = errors.New("outgoing wallet not found")
	ErrInsufficientFunds    = errors.New("insufficient funds")
	ErrTargetWalletNotFound = errors.New("target wallet not found")
	ErrNotFound             = errors.New("resource not found")
)
