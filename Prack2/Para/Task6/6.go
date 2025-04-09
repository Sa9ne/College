package main

// Библиотека для вывода в консоль, думаю можно
import "fmt"

func main() {
	var X, Y int

	fmt.Print("Введите координату X: ")
	fmt.Scan(&X)
	fmt.Print("Введите координату Y: ")
	fmt.Scan(&Y)

	if X > 0 && Y > 0 {
		fmt.Println("Первая координатная четверть")
	} else if X < 0 && Y > 0 {
		fmt.Println("Вторая координатная четверть")
	} else if X < 0 && Y < 0 {
		fmt.Println("Третья координатная четверть")
	} else {
		fmt.Println("Четвертая координатная четверть")
	}
}
