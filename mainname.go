package main

import (
	"fmt"
)

func main() {
	var name string // Объявляем переменную для имени

	fmt.Print("Как тебя зовут? ")
	fmt.Scanln(&name) // Считываем ввод пользователя

	fmt.Printf("Привет, %s! Удачи в изучении Go!\n", name)
}
