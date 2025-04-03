package main

import "fmt"

func main() {
	var Num, Sum int

	fmt.Print("Введите число: ")
	fmt.Scan(&Num)

	for Num > 0 {
		Sum += Num % 10
		Num /= 10
	}

	fmt.Println("Сумма чисел ровна:", Sum)
}
