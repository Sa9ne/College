package main

// Библиотека для вывода в консоль, думаю можно
import "fmt"

func main() {
	var FStor, SStor, TStor int

	fmt.Print("Введите длину первой стороны: ")
	fmt.Scan(&FStor)
	fmt.Print("Введите длину второй стороны: ")
	fmt.Scan(&SStor)
	fmt.Print("Введите длину третей стороны: ")
	fmt.Scan(&TStor)

	if FStor < (SStor+TStor) &&
		SStor < (FStor+TStor) &&
		TStor < (FStor+TStor) {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}
