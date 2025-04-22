package main

import (
	"fmt"
	"math/rand"
)

func MaxABS() {

}

func AvgABS() {

}

func ABSotl() {
}

func Random(Range, LowBorder, HighBorder int) []int {
	Mass := make([]int, Range)

	for i := 0; i < Range; i++ {
		Mass[i] = rand.Intn(HighBorder-LowBorder+1) + LowBorder
	}

	return Mass
}

func main() {
	// Ввод пользователя
	var Range, LowBorder, HighBorder int

	fmt.Print("Длина списка: ")
	fmt.Scan(&Range)
	fmt.Print("Нижняя граница генерации: ")
	fmt.Scan(&LowBorder)
	fmt.Print("Верхняя граница генерации: ")
	fmt.Scan(&HighBorder)

	// Генерация рандомных списков

	ListOne := Random(Range, LowBorder, HighBorder)
	ListTwo := Random(Range, LowBorder, HighBorder)
	ListThree := Random(Range, LowBorder, HighBorder)
	ListFour := Random(Range, LowBorder, HighBorder)
	fmt.Printf("Список 1: %v\nСписок 2: %v\nСписок 3: %v\nСписок 4: %v\n", ListOne, ListTwo, ListThree, ListFour)
}
