package main

import "fmt"

func main() {
	var Fchsl, Schsl int

	fmt.Println("Введите два числа:")
	fmt.Scan(&Fchsl)
	fmt.Scan(&Schsl)

	if Fchsl == Schsl {
		fmt.Println("Числа равны")
	}

	if Fchsl < Schsl {
		fmt.Printf("Меньшее число: %d", Fchsl)
	} else {
		fmt.Printf("Меньшее число: %d", Schsl)
	}
}
