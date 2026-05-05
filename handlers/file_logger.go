package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type LogRequest struct {
	Message string `json:"message"`
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Нужен POST запрос", http.StatusMethodNotAllowed)
		return
	}

	var req LogRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка в JSON", http.StatusBadRequest)
		return
	}

	// Открываем на дозапись (O_APPEND) или создаем (O_CREATE)
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Ошибка открытия файла", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(req.Message + "\n"); err != nil {
		http.Error(w, "Ошибка записи", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Строка успешно записана")
}
