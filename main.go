package main

import (
	"fmt"
	"my-first-go-app/handlers"
	"net/http"
)

func main() {
	router := handlers.SetupRoutes()
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
