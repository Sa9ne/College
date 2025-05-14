package main

import "fmt"

func main() {
	// Ввод пользователя
	var N int
	fmt.Print("Введите число N: ")
	fmt.Scan(&N)

	// Факториал
	sum := 0
	factorial := 1

	for i := 1; i <= N; i++ {
		factorial *= i
		sum += factorial
	}

	fmt.Println("Сумма факториалов: ", sum)
}
