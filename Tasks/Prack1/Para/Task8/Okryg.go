package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64

	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	fmt.Printf("Округление до наибольшего целого: %.0f\n", math.Ceil(num))
	fmt.Printf("Округление до наименьшего целого: %.0f\n", math.Floor(num))
}
