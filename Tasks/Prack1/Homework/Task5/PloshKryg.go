package main

import (
	"fmt"
	"math"
)

func main() {
	var R float64
	Pi := 3.1415926535

	fmt.Print("Введите радиус круга: ")
	fmt.Scan(&R)

	S := Pi * math.Pow(R, 2)
	fmt.Printf("Радиус круга равен: %f", S)
}
