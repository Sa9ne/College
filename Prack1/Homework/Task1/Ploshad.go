package main

import "fmt"

func main() {
	var FirstKat, SecKat float32

	fmt.Print("Введите длину катета: ")
	fmt.Scan(&FirstKat)
	fmt.Print("Введите длину катета: ")
	fmt.Scan(&SecKat)

	fmt.Println("Длинна площади равна:", (FirstKat*SecKat)/2)
}
