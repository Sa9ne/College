package main

import (
	"fmt"
)

func main() {
	var Rost, Ves float64

	fmt.Print("Введите вес человека (кг): ")
	fmt.Scan(&Ves)
	fmt.Print("Введите рост человека (см): ")
	fmt.Scan(&Rost)

	Rost /= 100
	IMT := Ves / (Rost * Rost)

	switch {
	case IMT <= 16:
		fmt.Println("Выраженный дефицит массы тела")
	case IMT > 16 && IMT <= 18.5:
		fmt.Println("Недостаточная масса тела")
	case IMT > 18.5 && IMT <= 25:
		fmt.Println("Норма")
	case IMT > 25 && IMT <= 30:
		fmt.Println("Избыточная масса тела")
	case IMT > 30 && IMT <= 35:
		fmt.Println("Ожирение 1 степени")
	case IMT > 35 && IMT <= 40:
		fmt.Println("Ожирение 2 степени")
	case IMT > 40:
		fmt.Println("Ожирение 3 степени")
	}
}
