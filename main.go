package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

// CalcRequest теперь содержит и теги для JSON, и методы для расчетов
type CalcRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

// Методы привязаны напрямую к структуре запроса
func (r *CalcRequest) Add() float64      { return float64(r.A + r.B) }
func (r *CalcRequest) Subtract() float64 { return float64(r.A - r.B) }
func (r *CalcRequest) Multiply() float64 { return float64(r.A * r.B) }
func (r *CalcRequest) Divide() (float64, error) {
	if r.B == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return float64(r.A) / float64(r.B), nil
}

type CalcResponse struct {
	Action string  `json:"action"`
	Result float64 `json:"result"`
}

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

	// Выбираем действие
	actions := []string{"сложение", "вычитание", "умножение", "деление"}
	action := actions[rand.Intn(len(actions))]

	var result float64
	var err error

	// Вызываем методы прямо у объекта запроса req
	switch action {
	case "сложение":
		result = req.Add()
	case "вычитание":
		result = req.Subtract()
	case "умножение":
		result = req.Multiply()
	case "деление":
		result, err = req.Divide()
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CalcResponse{
		Action: action,
		Result: result,
	})
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