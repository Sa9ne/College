package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Генерируем запросы
	for i := 0; i < 5; i++ {
		FNum := rand.Intn(10) + 1
		SNum := rand.Intn(10) + 1

		// + - / * %
		Operant := rand.Intn(6) + 1

		var CorrectAnswer int
		var UserAnswer int

		// Выводим пользователю и ждем ответа
		switch Operant {
		case 1:
			fmt.Printf("%d + %d = ?\n", FNum, SNum)
			CorrectAnswer = FNum + SNum
		case 2:
			fmt.Printf("%d - %d = ?\n", FNum, SNum)
			CorrectAnswer = FNum - SNum
		case 3:
			if SNum == 0 {
				SNum = rand.Intn(9) + 1 // Если 0, генерируем новое число
			}
			fmt.Printf("%d / %d = ?\n", FNum, SNum)
			CorrectAnswer = FNum / SNum
		case 4:
			fmt.Printf("%d * %d = ?\n", FNum, SNum)
			CorrectAnswer = FNum * SNum
		case 5:
			if SNum == 0 {
				SNum = rand.Intn(9) + 1 // Если 0, генерируем новое число
			}
			fmt.Print(FNum, " % ", SNum, "= ?\n")
			CorrectAnswer = FNum % SNum
		}
		// Проверяем ответ
		fmt.Scan(&UserAnswer)

		if UserAnswer == CorrectAnswer {
			fmt.Println("Правильный ответ")
		} else {
			fmt.Println("Не правильный ответ")
		}
	}
}
