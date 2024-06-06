// написаны все обработчки эндпоинтов
package handler

import "net/http"

// Обработчик для маршрута "/hello"
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Отправка приветственного сообщения
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Привет, мир"))
}
