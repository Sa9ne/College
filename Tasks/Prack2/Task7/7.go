package main

// Библиотека для вывода в консоль, думаю можно
import "fmt"

func main() {
	var NumMonth int
	fmt.Print("Введите номер месяца (1-12): ")
	fmt.Scan(&NumMonth)

	if NumMonth < 1 || NumMonth > 12 {
		fmt.Println("Ошибка: несуществующий номер месяца.")
		return
	}

	switch {
	case NumMonth == 12 || NumMonth == 1 || NumMonth == 2:
		fmt.Print("Время года: Зима")
	case NumMonth >= 3 && NumMonth <= 5:
		fmt.Print("Время года: Весна")
	case NumMonth >= 6 && NumMonth <= 8:
		fmt.Print("Время года: Лето")
	case NumMonth >= 9 && NumMonth <= 11:
		fmt.Print("Время года: Осень")
	}

	switch NumMonth {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Print(" 31")
	case 4, 6, 9, 11:
		fmt.Print(" 30")
	case 2:
		fmt.Print(" 28")
	}
}
