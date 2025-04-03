package main

import "fmt"

func main() {
	var Num int

	fmt.Print("Введите число, которое нужно перевернуть: ")
	fmt.Scan(&Num)

	LastNum := Num % 10
	FirstNum := Num / 100
	MiddleNum := (Num / 10) % 10

	Reversed := LastNum*100 + MiddleNum*10 + FirstNum

	fmt.Print("Перевернутое трехзначное число равно: ", Reversed)
}
