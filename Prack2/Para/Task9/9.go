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
	}

	if X+1 == A && Y+2 == B {
		fmt.Println("Yes")
	} else if X-1 == A && Y+2 == B {
		fmt.Println("Yes")
	} else if X+1 == A && Y-2 == B {
		fmt.Println("Yes")
	} else if X-1 == A && Y-2 == B {
		fmt.Println("Yes")
	} else if Y+1 == A && X+2 == B {
		fmt.Println("Yes")
	} else if Y-1 == A && X+2 == B {
		fmt.Println("Yes")
	} else if Y+1 == A && X-2 == B {
		fmt.Println("Yes")
	} else if Y-1 == A && X-2 == B {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
