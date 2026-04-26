package main

import (
	"encoding/json" // Пакет для работы с JSON
	"fmt"
	"math/rand"
	"net/http"
)

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

// Структуры данных
type CalcRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type CalcResponse struct {
	Action string  `json:"action"`
	Result float64 `json:"result"`
}

// --- ЛОГИКА ВЫЧИСЛЕНИЙ (Вынесена отдельно) ---
func performCalculation(a, b int) (string, float64, error) {
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
			return action, 0, fmt.Errorf("деление на ноль")
		}
		result = float64(a) / float64(b)
	}

	return action, result, nil
}

// --- ОБРАБОТЧИК ЗАПРОСА ---
func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Нужен POST запрос", http.StatusMethodNotAllowed)
		return
	}

	var req CalcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка в JSON", http.StatusBadRequest)
		return
	}

	// Вызываем функцию вычислений
	action, result, err := performCalculation(req.A, req.B)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := CalcResponse{
		Action: action,
		Result: result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/calculate", calculateHandler)
	return mux
}

func main() {
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", setupRoutes())
}