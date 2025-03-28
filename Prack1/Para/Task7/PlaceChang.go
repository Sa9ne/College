package main

import "fmt"

func main() {
	var A, B float32

	fmt.Print("Введите значение переменной A: ")
	fmt.Scan(&A)
	fmt.Print("Введите значение переменной B: ")
	fmt.Scan(&B)

	fmt.Printf("A = %.1f", B)
	fmt.Printf("\nB = %.1f", A)
}
