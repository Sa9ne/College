package main

import "fmt"

func main() {
	var Lis, Cat, Money int

	fmt.Print("Введите количество лисиц: ")
	fmt.Scan(&Lis)

	fmt.Print("Введите количество котов: ")
	fmt.Scan(&Cat)

	fmt.Print("Введите количество монет: ")
	fmt.Scan(&Money)

	TotalAnimal := Lis + Cat

	if TotalAnimal == 0 {
		fmt.Println("Животных вообще нет")
		return
	}

	CoinPerAnimal := Money / TotalAnimal
	CoinShapka := Money % TotalAnimal

	fmt.Println("\nКаждому достанется монет:", CoinPerAnimal)
	fmt.Println("Останется в шляпе монет:", CoinShapka)
}
