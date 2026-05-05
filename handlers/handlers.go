package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Нужен POST запрос", http.StatusMethodNotAllowed)
		return
	}

	var req CalcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка в JSON", http.StatusBadRequest)
		return
	}

	actions := []string{"сложение", "вычитание", "умножение", "деление"}
	action := actions[rand.Intn(len(actions))]

	var result float64
	var err error

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
	json.NewEncoder(w).Encode(CalcResponse{Action: action, Result: result})
}

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/calculate", CalculateHandler)
	mux.HandleFunc("/log", FileHandler) // Новый хэндлер
	return mux
}
