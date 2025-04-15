package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Граница генерации чисел
	var min, max int

	fmt.Print("Введите нижнюю границу генерации: ")
	fmt.Scan(&min)
	fmt.Print("Введите верхнюю границу генерации: ")
	fmt.Scan(&max)
	// Генерация числа
	Random := rand.Intn(max-min+1) + min
	// Проверка угадал ли пользователь
	NumCheck := false
	NumFound := 0

	for !NumCheck {
		fmt.Print("Угадайте число: ")
		fmt.Scan(&NumFound)

		if NumFound > max || NumFound < min {
			fmt.Println("Вы ввели значение больше/меньше границы")
			NumCheck = false
		} else if NumFound == Random {
			NumCheck = true
		} else {
			fmt.Println("Не угадал")
			NumCheck = false
		}
	}

	// Вывод
	fmt.Println("Молодец, число и вправду: ", NumFound)
}
