package main

import "fmt"

func DeletedNum(DeleteNum int) {

}

func main() {
	// Ввод числа
	var Num, DeleteNum int

	fmt.Print("Введите число: ")
	fmt.Scan(&Num)
	fmt.Print("Введите цифру для удаления: ")
	fmt.Scan(&DeleteNum)

	// Вызов функции и результат
	DeletedNum(DeleteNum)

}
