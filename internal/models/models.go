package models

import "time"

type Wallet struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type Transaction struct {
	ID         string    `json:"id"`
	Time       time.Time `json:"time"`
	FromWallet string    `json:"from_wallet"`
	ToWallet   string    `json:"to_wallet"`
	Amount     float64   `json:"amount"`
}
