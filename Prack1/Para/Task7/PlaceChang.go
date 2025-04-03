package main

import "fmt"

func main() {
	var A, B float32

	fmt.Print("Введите значение переменной A: ")
	fmt.Scan(&A)
	fmt.Print("Введите значение переменной B: ")
	fmt.Scan(&B)

	// Первый способ
	A = B
	B = A

	fmt.Printf("A = %.1f", A)
	fmt.Printf("\nB = %.1f", B)

	// Второй способ
	NewPeremen := B
	B = A
	A = NewPeremen

	fmt.Printf("\nB = %.1f", B)
	fmt.Printf("\n\nA = %.1f", A)

	// Третий способ
	A = A + B
	B = A - B
	A = A - B

	fmt.Printf("\n\nA = %.1f", A)
	fmt.Printf("\nB = %.1f", B)
}
