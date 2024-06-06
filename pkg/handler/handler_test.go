package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder() // тип Recorder удовлетовряет интефрейсу ResponseWriter
	handler := http.HandlerFunc(HelloHandler)

	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Привет, мир" // когда буду API тестировать здесь будет уже объект json , к примеру, expected := `{"alive": true}`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
