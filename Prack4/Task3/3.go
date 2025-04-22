package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Максимальное абсолютное отклонение
func MaxABS(a, b []int) int {
	max := 0
	for i := range a {
		diff := int(math.Abs(float64(a[i] - b[i])))
		if diff > max {
			max = diff
		}
	}
	return max
}

// Среднее абсолютное отклонение
func AvgABS(a, b []int) float64 {
	sum := 0
	for i := range a {
		sum += int(math.Abs(float64(a[i] - b[i])))
	}
	return float64(sum) / float64(len(a))
}

// Абсолютное отклонение между двумя списками
func ABSotl(a, b []int) []int {
	result := make([]int, len(a))
	for i := range a {
		result[i] = int(math.Abs(float64(a[i] - b[i])))
	}
	return result
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

	// Вывод отклонений
	fmt.Printf("\nМакс. отклонение (1 и 2): %d\n", MaxABS(ListOne, ListTwo))
	fmt.Printf("Среднее отклонение (1 и 2): %.2f\n", AvgABS(ListOne, ListTwo))
	fmt.Printf("Макс. отклонение (1 и 3): %d\n", MaxABS(ListOne, ListThree))
	fmt.Printf("Среднее отклонение (1 и 3): %.2f\n", AvgABS(ListOne, ListThree))
	fmt.Printf("Макс. отклонение (1 и 4): %d\n", MaxABS(ListOne, ListFour))
	fmt.Printf("Среднее отклонение (1 и 4): %.2f\n", AvgABS(ListOne, ListFour))
}
