package main

// Библиотека для вывода в консоль, думаю можно
import "fmt"

func main() {
	var M, N, K int

	fmt.Print("M = ")
	fmt.Scan(&M)
	fmt.Print("N = ")
	fmt.Scan(&N)
	fmt.Print("K = ")
	fmt.Scan(&K)

	if K%M == 0 || N%K == 0 {
		fmt.Print("Можно")
	} else {
		fmt.Print("Нельзя")
	}
}
