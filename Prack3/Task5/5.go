package main

import "fmt"

func main() {
	// Ввод пользователя
	var Length, Element int
	var Mas []int

	fmt.Print("Введите длину списка: ")
	fmt.Scan(&Length)
	fmt.Println("Введите элементы списка: ")
	for i := 0; i < Length; i++ {
		fmt.Scan(&Element)
		Mas = append(Mas, Element)
	}

	// Проверяем список на повторные элементы
	var DoubleElement []int

	for i := 0; i < Length; i++ {
		// Хотелось бы бинарный поиск, но сортировку нельзя((
		count := 0

		for j := 0; j < Length; j++ {
			if Mas[i] == Mas[j] {
				count++
			}
		}

		if count == 1 {
			DoubleElement = append(DoubleElement, Mas[i])
		}
	}

	// Выводим данные
	fmt.Println(DoubleElement)
}
