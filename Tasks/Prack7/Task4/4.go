package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n, m, Place int

	fmt.Print("Введите количество рядов: ")
	fmt.Scan(&n)
	fmt.Print("Введите количество мест в ряду: ")
	fmt.Scan(&m)

	// Создаем двумерный массив зала
	// Создаем зал в котором будет n кол-во рядов
	Hall := make([][]int, n)

	for i := range Hall {
		// Создаем кол-во мест в ряду
		Hall[i] = make([]int, m)
		for j := range Hall[i] {
			// Задаем рандом
			// 70% мест будут свободны, 30% — заняты
			if rand.Intn(100) < 30 {
				Hall[i][j] = 1
			}
		}
	}

	// Вывод зала
	fmt.Println("\nСхема зала:")
	for i := 0; i < n; i++ {
		// Вывод кол-во рядов
		fmt.Printf("Ряд %2d: ", i+1)
		for j := 0; j < m; j++ {
			// Вывод занятой и не занятой позиции
			// Идет по i и j, так как это двумерный массив
			fmt.Print(Hall[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Print("\nВведите количество билетов, которые нужно купить подряд: ")
	fmt.Scan(&Place)

	// Поиск подходящих рядов
	Found := false

	for i := 0; i < n; i++ {
		// Свободное место
		count := 0
		// Найдено в ряде
		rowFound := false
		for j := 0; j < m; j++ {
			if Hall[i][j] == 0 {
				// Считаем свободные места
				count++
				if count == Place && !rowFound {
					fmt.Printf("Ряд %d\n", i+1)
					Found = true
					rowFound = true
				}
			} else {
				count = 0
			}
		}
	}

	if !Found {
		fmt.Println("Нет места")
	}
}
