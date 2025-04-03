package main

import (
	"fmt"
	"math"
)

func main() {
	var FirsClass, SecClass, ThirdClass int
	var FirstWorkTable, SecWorkTable, ThiWorkTable float64

	fmt.Print("В первом классе: ")
	fmt.Scan(&FirsClass)
	fmt.Print("В первом классе: ")
	fmt.Scan(&SecClass)
	fmt.Print("В первом классе: ")
	fmt.Scan(&ThirdClass)

	FirstWorkTable = float64(FirsClass) / 2
	SecWorkTable = float64(SecClass) / 2
	ThiWorkTable = float64(ThirdClass) / 2

	AllWorkTable := math.Ceil(FirstWorkTable) + math.Ceil(SecWorkTable) + math.Ceil(ThiWorkTable)

	fmt.Printf("Нужно %.0f парт(ы)", AllWorkTable)
}
