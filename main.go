package main

import (
	"fmt"
	"net/http"
)

// setupRoutes теперь содержит несколько эндпоинтов
func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Главная страница
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Добро пожаловать на главную страницу!")
	})

	// Эндпоинт /hello
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Привет! Ты вызвал эндпоинт Hello.")
	})

	// Эндпоинт /about
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Это сервер на Go. Версия 1.0")
	})

	return mux
}

func main() {
	router := setupRoutes()
	fmt.Println("Сервер запущен. Проверь адреса:")
	fmt.Println("http://localhost:8080/")
	fmt.Println("http://localhost:8080/hello")
	fmt.Println("http://localhost:8080/about")

	http.ListenAndServe(":8080", router)
}
