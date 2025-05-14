package main

import (
	"fmt"
)

func main() {
	// Ввод данных от пользователя
	var Vklad, Prots float64
	var months int

	fmt.Print("Сумма вклада: ")
	fmt.Scan(&Vklad)
	fmt.Print("Процентная ставка (проценты в год): ")
	fmt.Scan(&Prots)
	fmt.Print("Срок вклада (месяцы): ")
	fmt.Scan(&months)

	// Для каждого месяца рассчитываем сумму с процентами
	for month := 1; month <= months; month++ {
		Vklad *= (1 + Prots/100/12)
		fmt.Printf("%d месяц: %f\n", month, Vklad)
	}
}
