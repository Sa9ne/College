package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var length int
	var lower, upper float64

	// Ввод от пользователя
	fmt.Print("Длина списка: ")
	fmt.Scan(&length)
	fmt.Print("Нижняя граница генерации: ")
	fmt.Scan(&lower)
	fmt.Print("Верхняя граница генерации: ")
	fmt.Scan(&upper)

	rand.Seed(time.Now().UnixNano())

	numbers := make([]float64, length)

	for i := 0; i < length; i++ {
		if rand.Intn(2) == 0 {
			numbers[i] = float64(rand.Intn(int(upper)-int(lower)+1) + int(lower))
		} else {
			numbers[i] = lower + rand.Float64()*(upper-lower)
		}
	}

	fmt.Printf("Сгенерированный список: %v\n", numbers)

	var sum float64
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("Сумма чисел: %.2f\n", sum)

	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Printf("Максимальное число в списке: %.2f\n", max)

	var evens []int
	for _, num := range numbers {
		if float64(int(num)) == num && int(num)%2 == 0 {
			evens = append(evens, int(num))
		}
	}
	fmt.Printf("Чётные числа: %v\n", evens)
}
