package main

import "fmt"

func Kalkyliator() {
	var a, b int

	fmt.Println("Введите первое число")
	fmt.Scan(&a)
	fmt.Println("Введите второе число")
	fmt.Scan(&b)

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	if b != 0 {
		fmt.Printf("%d * %d = %.1f\n", a, b, float32(a)/float32(b))
		fmt.Printf("%d // %d = %d\n", a, b, a/b)
		fmt.Printf("%d %% %d = %d\n", a, b, a%b)
	} else {
		fmt.Println("Делить на 0 нельзя")
	}
}
