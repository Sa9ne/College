package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Sort(Mas []int) {
	MasRange := len(Mas)
	// Смотрим длину массива и вычетам один т.к. массив начинается в 0
	for i := 0; i < MasRange-1; i++ {
		// Делаем второй вложенный цикл
		for j := 0; j < MasRange-i-1; j++ {
			if Mas[j] > Mas[j+1] {
				// Меняем элементы местами
				Mas[j], Mas[j+1] = Mas[j+1], Mas[j]
			}
		}
	}
	// Хочу глянуть сам массив
	fmt.Println(Mas)
}

func main() {
	var Lower, Upper int
	fmt.Println("Введите нижнюю границу")
	fmt.Scan(&Lower)

	fmt.Println("Введите верхнюю границу")
	fmt.Scan(&Upper)

	for range 3 {
		// Генерация рандома и занесения в массив
		var Mas []int

		// Взял 20 как длину массива, не очень понял что именно нужно взять
		for i := 0; i < 20; i++ {
			Random := rand.Intn(Upper-Lower+1) + Lower
			Mas = append(Mas, Random)
		}

		// Засекаем время
		TimeStart := time.Now()
		Sort(Mas)
		Duration := time.Since(TimeStart)

		fmt.Println("Прошло времени с момента начало", Duration)
	}
}
