package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		http.Error(w, "oops", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Сервер запущен на http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
