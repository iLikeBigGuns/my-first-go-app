package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

// Calculator — структура, которая будет содержать наши числа и методы
type Calculator struct {
	A int
	B int
}

// Методы для каждого типа вычисления
func (c Calculator) Add() float64      { return float64(c.A + c.B) }
func (c Calculator) Subtract() float64 { return float64(c.A - c.B) }
func (c Calculator) Multiply() float64 { return float64(c.A * c.B) }
func (c Calculator) Divide() (float64, error) {
	if c.B == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return float64(c.A) / float64(c.B), nil
}

// Структуры для JSON
type CalcRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type CalcResponse struct {
	Action string  `json:"action"`
	Result float64 `json:"result"`
}

// --- ЛОГИКА ВЫБОРА МЕТОДА ---
func performCalculation(a, b int) (string, float64, error) {
	calc := Calculator{A: a, B: b}
	
	actions := []string{"сложение", "вычитание", "умножение", "деление"}
	action := actions[rand.Intn(len(actions))]
	
	var result float64
	var err error

	// Вызываем конкретный метод в зависимости от выбранного действия
	switch action {
	case "сложение":
		result = calc.Add()
	case "вычитание":
		result = calc.Subtract()
	case "умножение":
		result = calc.Multiply()
	case "деление":
		result, err = calc.Divide()
	}

	return action, result, err
}

// --- ОБРАБОТЧИК ---
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

	action, result, err := performCalculation(req.A, req.B)
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
