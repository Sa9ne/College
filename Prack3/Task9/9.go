package main

import "fmt"

func main() {
	// Ввод пользователя
	var N int
	fmt.Print("Введите число N: ")
	fmt.Scan(&N)

	// Проверка всех остальных чисел на простоту до N
	if N < 2 {
		fmt.Println("Простых чисел нет")
	}

	for num := 2; num <= N; num++ {
		isPrime := true

		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
			}
		}

		if isPrime {
			fmt.Print(num, " ")
		}
	}
}
