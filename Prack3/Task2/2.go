package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Ввод пользователя
	var Range, Min, Max int

	fmt.Print("Длинна списка: ")
	fmt.Scan(&Range)
	fmt.Print("Нижняя граница генерации: ")
	fmt.Scan(&Min)
	fmt.Print("Верхняя граница генерации: ")
	fmt.Scan(&Max)

	// Генерация рандома и занесения в массив
	var Mas []int

	for i := 0; i < Range; i++ {
		Random := rand.Intn(Max-Min+1) + Min
		Mas = append(Mas, Random)
	}

	// Поиск максимального значения
	MaxNum := Mas[0]

	for i := 1; i < Range; i++ {
		if Mas[i] > MaxNum {
			MaxNum = Mas[i]
		}
	}

	// Вывод информации
	fmt.Println("Сгенерированный список: ", Mas)
	fmt.Println("Максимальное значение: ", MaxNum)
}
