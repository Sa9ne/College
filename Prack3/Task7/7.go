package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Имитация 1000 бросков
	Stat := make([]int, 6)

	for i := 0; i < 1000; i++ {
		// Генерация выпадения кубика
		Random := rand.Intn(6) + 1
		// Подсчет выпавшего числа
		Stat[Random-1]++
	}

	// Вывод результатов в консоль
	for i := 0; i < 6; i++ {
		fmt.Printf("Значение %d - %d\n", i+1, Stat[i])
	}
}
