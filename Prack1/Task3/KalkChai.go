package main

import "fmt"

func main() {
	var Money float32

	fmt.Print("Сумма вашего чека: ")
	fmt.Scan(&Money)

	fmt.Println("10 % =", Money*0.1)
	fmt.Println("10 % =", Money*0.15)
	fmt.Println("10 % =", Money*0.2)
}
