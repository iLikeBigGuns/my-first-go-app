package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Регистрируем эндпоинт "/" и функцию-обработчик
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Привет! Это твой первый HTTP сервер на Go.")
	})

	fmt.Println("Сервер запущен на http://localhost:8080")
	
	// Запускаем сервер на порту 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска:", err)
	}
}
