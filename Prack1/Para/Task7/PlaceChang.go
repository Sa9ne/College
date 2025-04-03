package main

import "fmt"

func main() {
	var A, B float32

	fmt.Print("Введите значение переменной A: ")
	fmt.Scan(&A)
	fmt.Print("Введите значение переменной B: ")
	fmt.Scan(&B)

	// Первый способ
	fmt.Printf("A = %.1f", B)
	fmt.Printf("\nB = %.1f", A)

	// Второй способ
	NewPeremen := B
	fmt.Printf("\n\nA = %.1f", NewPeremen)
	NewPeremen = A
	fmt.Printf("\nB = %.1f", NewPeremen)

	// Третий способ
	A = A + B
	B = A - B
	A = A - B

	fmt.Printf("\n\nA = %.1f", A)
	fmt.Printf("\nB = %.1f", B)
}
