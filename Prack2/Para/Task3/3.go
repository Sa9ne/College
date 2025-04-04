package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64

	fmt.Print("Введите коэффициент A: ")
	fmt.Scan(&A)
	fmt.Print("Введите коэффициент B: ")
	fmt.Scan(&B)
	fmt.Print("Введите свободный C: ")
	fmt.Scan(&C)

	Disck := B*B - 4*A*C
	fmt.Printf("Дискриминант равен: %.2f", Disck)

	if Disck > 0 {
		x1 := (-B + math.Sqrt(Disck)/(2*A))
		x2 := (-B - math.Sqrt(Disck)/(2*A))
		fmt.Printf("\nДва корня:%2.F и%2.F", x1, x2)
	} else if Disck == 0 {
		x := -B / (2 * A)
		fmt.Printf("\nКорень равен: %2.F", x)
	} else {
		fmt.Println("Корней нет")
	}
}
