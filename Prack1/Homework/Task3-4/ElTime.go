package main

import "fmt"

func main() {
	var Time int

	fmt.Print("Введите число, которое преобразуется во время: ")
	fmt.Scan(&Time)

	if Time < 1 {
		fmt.Println("Введите число больше 0")
		return
	}

	day := Time / 86400
	hours := (Time & 86400) / 3600
	min := (Time % 3600) / 60
	sec := Time % 60

	if day > 0 {
		fmt.Printf("%d дней\n", day)
	}
	if hours > 0 {
		fmt.Printf("%d часов\n", hours)
	}
	if min > 0 {
		fmt.Printf("%d минут\n", min)
	}
	if sec > 0 || Time == 0 {
		fmt.Printf("%d секунд\n", sec)
	}

	fmt.Printf("\n%02d:%02d:%02d:%02d", day, hours, min, sec)
}
