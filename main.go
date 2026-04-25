package main

import (
	"encoding/json" // Пакет для работы с JSON
	"fmt"
	"math/rand"
	"net/http"
)

// Описываем структуру входящих данных
type CalcRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// GET: Главная страница
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Добро пожаловать на главную страницу!")
	})

	// GET: Эндпоинт /hello
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Привет! Ты вызвал эндпоинт Hello.")
	})	
	
	// POST: Связываем путь с функцией calculateHandler
	mux.HandleFunc("/calculate", calculateHandler)

	return mux
}

// POST: Эндпоинт /calculate
func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Используйте POST запрос", http.StatusMethodNotAllowed)
		return
	}

	// 1. Создаем переменную типа нашей структуры
	var req CalcRequest

	// 2. Читаем JSON из тела запроса и записываем в переменную req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Ошибка: неверный формат JSON", http.StatusBadRequest)
		return
	}

	// 3. Используем данные из структуры (теперь это уже числа, а не строки!)
	a := req.A
	b := req.B

	actions := []string{"сложение", "умножение", "деление", "вычитание"}
	action := actions[rand.Intn(len(actions))]
	var result float64

	switch action {
	case "сложение":
		result = float64(a + b)
	case "вычитание":
		result = float64(a - b)
	case "умножение":
		result = float64(a * b)
	case "деление":
		if b == 0 {
			fmt.Fprintln(w, "На ноль делить нельзя")
			return
		}
		result = float64(a) / float64(b)
	}

	fmt.Fprintf(w, "Действие: %s. Результат: %.2f\n", action, result)
}

func main() {
	router := setupRoutes()
	fmt.Println("Сервер запущен. Ожидаю JSON на http://localhost:8080/calculate")
	http.ListenAndServe(":8080", router)
}