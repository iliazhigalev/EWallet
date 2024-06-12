package handler

import (
	"database/sql"
	"encoding/json"
	"ewallet/database"
	"ewallet/models"
	"ewallet/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var mockDB *sql.DB
var mock sqlmock.Sqlmock

func init() {
	var err error
	mockDB, mock, err = sqlmock.New()
	if err != nil {
		panic("failed to open sqlmock database connection")
	}
}

func TestCreateWallet(t *testing.T) {

	utils.GenerateID = func() string {
		return "test-id"
	}

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/wallet", nil)

	oldConnectDb := database.ConnectDb
	defer func() { database.ConnectDb = oldConnectDb }()
	database.ConnectDb = func() (*sql.DB, error) {
		return mockDB, nil
	}

	mock.ExpectExec("INSERT INTO wallets").
		WithArgs("test-id", 100.0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	CreateWallet(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var wallet models.Wallet
	err := json.NewDecoder(recorder.Body).Decode(&wallet)
	assert.NoError(t, err)
	assert.Equal(t, "test-id", wallet.ID)
	assert.Equal(t, 100.0, wallet.Balance)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
