package main

import "fmt"

func main() {
	var a, b, c float32

	fmt.Println("Введите три любых числа")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	fmt.Println("Среднее арифметическое =", (a+b+c)/3)
}
