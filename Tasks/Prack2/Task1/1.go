package main

import "fmt"

func main() {
	var A int

	fmt.Print("Введите целое число: ")
	fmt.Scan(&A)

	Chisl := A % 2
	if Chisl > 0 {
		fmt.Println("Число не четное")
	} else {
		fmt.Println("Число четное")
	}
}
