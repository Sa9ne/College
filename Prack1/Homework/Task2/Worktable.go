package main

import (
	"fmt"
	"math"
)

func main() {
	var FirsClass, SecClass, ThirdClass int
	var AllTable float64

	fmt.Print("В первом классе: ")
	fmt.Scan(&FirsClass)
	fmt.Print("В первом классе: ")
	fmt.Scan(&SecClass)
	fmt.Print("В первом классе: ")
	fmt.Scan(&ThirdClass)

	AllTable = float64(FirsClass+SecClass+ThirdClass) / 2

	fmt.Printf("Нужно %.0f парт(ы)", math.Ceil(AllTable))
}
