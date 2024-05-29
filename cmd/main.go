package main

// из файла main.go происходит запуск прогаммы
import (
	"fmt"
	"log"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server started")

}
func main() {

	log.Println("Server started at :3000")
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":3000", nil)
}
