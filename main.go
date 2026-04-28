package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := setupRoutes()
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", router)
}