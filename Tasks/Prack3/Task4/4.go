package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Ввод пользователя
	var Length int

	fmt.Print("Введите длину пароля: ")
	fmt.Scan(&Length)

	// Создание всех символов и их объединение
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	special := "!@#$%^&*()-_=+[]{}<>?/|"

	AllSymbols := lower + upper + digits + special

	// Генерация пароля
	var Password string

	for i := 0; i < Length; i++ {
		rand := rand.Intn(len(AllSymbols))
		Password += string(AllSymbols[rand])
	}

	// Вывод пароля пользователю
	fmt.Println("Пароль: ", Password)
}
