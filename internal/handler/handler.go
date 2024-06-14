package handler

import (
	"database/sql"
	"encoding/json"
	"ewallet/internal/database"
	"ewallet/internal/models"
	"ewallet/pkg/utils"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	id := utils.GenerateID()
	balance := 100.0
	adminID := "admin1"

	db, err := database.ConnectDb()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Failed to create transaction", http.StatusInternalServerError)
		return
	}

	query := `INSERT INTO wallets (id, balance) VALUES ($1, $2)`
	_, err = tx.Exec(query, id, balance)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to create wallet", http.StatusInternalServerError)
		return
	}

	transactionQuery := `INSERT INTO transactions (id, time, from_wallet, to_wallet, amount) VALUES ($1, $2, $3, $4, $5)`
	transactionID := utils.GenerateID()
	_, err = tx.Exec(transactionQuery, transactionID, time.Now(), adminID, id, balance)
	if err != nil {
		tx.Rollback()
		log.Printf("Failed to log initial transaction: %v", err)
		http.Error(w, "Failed to log initial transaction", http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	response := models.Wallet{ID: id, Balance: balance}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendMoney(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fromWalletID := vars["walletId"]

	var request struct {
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	toWalletID := request.To
	amount := request.Amount

	db, err := database.ConnectDb()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var fromWalletBalance float64
	err = db.QueryRow("SELECT balance FROM wallets WHERE id = $1", fromWalletID).Scan(&fromWalletBalance)
	if err == sql.ErrNoRows {
		http.Error(w, "From wallet not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to retrieve from wallet", http.StatusInternalServerError)
		return
	}

	if fromWalletBalance < amount {
		http.Error(w, "Insufficient funds in from wallet", http.StatusBadRequest)
		return
	}

	var toWalletBalance float64
	err = db.QueryRow("SELECT balance FROM wallets WHERE id = $1", toWalletID).Scan(&toWalletBalance)
	if err == sql.ErrNoRows {
		http.Error(w, "To wallet not found", http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, "Failed to retrieve to wallet", http.StatusInternalServerError)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Failed to begin transaction", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("UPDATE wallets SET balance = balance - $1 WHERE id = $2", amount, fromWalletID)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to update from wallet balance", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("UPDATE wallets SET balance = balance + $1 WHERE id = $2", amount, toWalletID)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to update to wallet balance", http.StatusInternalServerError)
		return
	}

	transactionID := utils.GenerateID()
	_, err = tx.Exec("INSERT INTO transactions (id, time, from_wallet, to_wallet, amount) VALUES ($1, $2, $3, $4, $5)",
		transactionID, time.Now(), fromWalletID, toWalletID, amount)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to log transaction", http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetTransactionHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletID := vars["walletId"]

	db, err := database.ConnectDb()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var wallet models.Wallet
	err = db.QueryRow("SELECT id FROM wallets WHERE id = $1", walletID).Scan(&wallet.ID)
	if err == sql.ErrNoRows {
		log.Printf("Wallet not found: %v", walletID)
		http.Error(w, "Wallet not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Error querying wallet: %v", err)
		http.Error(w, "Failed to retrieve wallet", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT id, time, from_wallet, to_wallet, amount FROM transactions WHERE from_wallet = $1 OR to_wallet = $1", walletID)
	if err != nil {
		log.Printf("Error querying transactions: %v", err)
		http.Error(w, "Failed to retrieve transactions", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.Time, &transaction.FromWallet, &transaction.ToWallet, &transaction.Amount); err != nil {
			log.Printf("Failed to scan transaction: %v", err)
			http.Error(w, "Failed to scan transaction", http.StatusInternalServerError)
			return
		}
		transactions = append(transactions, transaction)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v", err)
		http.Error(w, "Error iterating through transactions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(transactions); err != nil {
		log.Printf("Failed to encode transactions: %v", err)
		http.Error(w, "Failed to encode transactions", http.StatusInternalServerError)
	}
}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletId := vars["walletId"]

	db, err := database.ConnectDb()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var wallet models.Wallet
	err = db.QueryRow(`SELECT id, balance FROM wallets WHERE id=$1`, walletId).Scan(&wallet.ID, &wallet.Balance)
	if err != nil {
		http.Error(w, "Wallet not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wallet)
}
