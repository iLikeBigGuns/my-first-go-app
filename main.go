package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
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

	// POST: Эндпоинт /calculate
	mux.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		// 1. Проверяем, что это POST запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Используйте POST запрос", http.StatusMethodNotAllowed)
			return
		}

		// 2. Получаем числа из параметров URL
		val1 := r.URL.Query().Get("a")
		val2 := r.URL.Query().Get("b")

		// 3. Конвертируем текст в целые числа
		a, err1 := strconv.Atoi(val1)
		b, err2 := strconv.Atoi(val2)

		if err1 != nil || err2 != nil {
			fmt.Fprintln(w, "Ошибка: укажите целые числа в параметрах a и b")
			return
		}

		// 4. Случайное действие
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
	})

	return mux
}

func main() {
	router := setupRoutes()
	fmt.Println("Сервер запущен. POST эндпоинт: http://localhost:8080/calculate?a=10&b=5")

	http.ListenAndServe(":8080", router)
}