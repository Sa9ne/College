package main

// Библиотека для вывода в консоль, думаю можно
import "fmt"

func main() {
	var X, Y, A, B int

	fmt.Print("X: ")
	fmt.Scan(&X)
	fmt.Print("Y: ")
	fmt.Scan(&Y)
	fmt.Print("A: ")
	fmt.Scan(&A)
	fmt.Print("B: ")
	fmt.Scan(&B)

	if A < 0 || B < 0 {
		fmt.Println("No")
		return
	}

	RaznX := A - X
	RaznY := B - Y

	if (RaznX*RaznX == 1 && RaznY*RaznY == 4) || (RaznX*RaznX == 4 && RaznY*RaznY == 1) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
