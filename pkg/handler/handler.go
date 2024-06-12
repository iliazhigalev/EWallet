package handler

import (
	"encoding/json"
	"ewallet/database"
	"ewallet/models"
	"ewallet/utils"
	"net/http"
)

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	id := utils.GenerateID()
	balance := 100.0

	db, err := database.ConnectDb()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := `INSERT INTO wallets (id, balance) VALUES ($1, $2)`
	_, err = db.Exec(query, id, balance)
	if err != nil {
		http.Error(w, "Failed to create wallet", http.StatusInternalServerError)
		return
	}

	response := models.Wallet{ID: id, Balance: balance}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendMoney(w http.ResponseWriter, r *http.Request) {

}

func GetTransactionHistory(w http.ResponseWriter, r *http.Request) {

}

func GetWallet(w http.ResponseWriter, r *http.Request) {

}
