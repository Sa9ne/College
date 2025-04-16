package main

import (
	"fmt"
	"math"
)

func main() {
	// Занесения чисел в массив
	StopCycle := false
	Num := 0
	List := []int{}

	for !StopCycle {
		fmt.Print("Введите число: ")
		fmt.Scan(&Num)
		if Num == 0 {
			StopCycle = true
		} else {
			List = append(List, Num)
		}
	}

	// Счет Ср.Ариф
	var AVG float64

	for i := 0; i < len(List); i++ {
		AVG += float64(List[i])
	}

	AVG /= float64(len(List))

	// Поиск числа, который ближе к сред ариф.
	Closest := List[0]
	minDiff := math.Abs(float64(Closest) - AVG)

	for i := 0; i < len(List); i++ {
		val := List[i]
		diff := math.Abs(float64(val) - AVG)
		if diff < minDiff {
			minDiff = diff
			Closest = val
		}
	}

	// Вывод информации
	fmt.Println("Заданный список:", List)
	fmt.Println("Среднее арифметическое:", AVG)
	fmt.Println("Ближайшее значение:", Closest)
}
