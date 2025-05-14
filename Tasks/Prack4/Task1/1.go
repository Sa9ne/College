package main

import "fmt"

func DeletedNum(Num int, DeletedNum int) {
	GoodNum := 0

	// Алгоритм для уничтожения ненужного числа
	for Num > 0 {
		if Num%10 != DeletedNum {
			// Проверка на первый элемент
			if GoodNum > 0 {
				GoodNum *= 10
				GoodNum += Num % 10
			} else {
				GoodNum = Num % 10
			}
			Num /= 10
		} else {
			Num /= 10
		}
	}

	// Так как у нас получился перевернутый элемент, переворачиваем его
	result := 0

	for GoodNum != 0 {
		result = result*10 + GoodNum%10
		GoodNum /= 10
	}

	fmt.Println(result)
}

func main() {
	// Ввод числа
	var Num, DeleteNum int

	fmt.Print("Введите число: ")
	fmt.Scan(&Num)
	fmt.Print("Введите цифру для удаления: ")
	fmt.Scan(&DeleteNum)

	if DeleteNum/10 != 0 {
		fmt.Println("Что то много чисел для удаления")
		return
	}

	// Вызов функции и результат
	DeletedNum(Num, DeleteNum)

}
