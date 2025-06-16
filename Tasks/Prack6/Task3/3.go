package main

import (
	moduls "college/Prack6/Task3/Moduls"
	"fmt"
)

func main() {
	var Lower, Upper, Posled, FoundNum int

	fmt.Print("Введите нижнюю границу генерации: ")
	fmt.Scan(&Lower)
	fmt.Print("Введите верхнюю границу генерации: ")
	fmt.Scan(&Upper)
	fmt.Print("Введите последовательность: ")
	fmt.Scan(&Posled)

	Mas := moduls.Random(Upper, Lower, Posled)
	fmt.Println(Mas)

	fmt.Print("Введите число которое вы хотите найти: ")
	fmt.Scan(&FoundNum)

	Index := moduls.BinarySearch(Mas, FoundNum)
	if Index == -1 {
		fmt.Println("Число не было найдено")
	}
}
