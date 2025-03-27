package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Введите первое число")
	fmt.Scan(&a)
	fmt.Println("Введите второе число")
	fmt.Scan(&b)

	fmt.Println("Сумма чисел ровна: ", a+b)
	fmt.Println("Разность чисел ровна: ", a-b)
	fmt.Println("Произведение чисел равно: ", a*b)
	if b != 0 {
		fmt.Println("Частное двух чисел равно: ", float32(a/b))
		fmt.Println("Целочисленное деление двух чисел равно: ", a/b)
		fmt.Println("Остаток от деления двух чисел равен: ", a%b)
	} else {
		fmt.Println("Делить на 0 нельзя")
	}
}
