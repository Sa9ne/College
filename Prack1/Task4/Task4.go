package main

import "fmt"

func main() {
	var FiveNum, Sum int

	fmt.Print("Введите пятизначное число: ")
	fmt.Scan(&FiveNum)

	if (FiveNum/10000) < 1 || (FiveNum/10000) > 10 {
		fmt.Println("Вы ввели не пятизначное число")
		return
	}

	for i := 0; i < 5; i++ {
		Sum += FiveNum % 10
		FiveNum /= 10
	}

	fmt.Println("Сумма чисел ровна: ", Sum)
}
