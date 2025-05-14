package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	Data, err := os.ReadFile("text.txt")
	if err != nil {
		panic(fmt.Sprintln("При чтении файла нашлась ошибка: ", err))
	}

	words := strings.Fields(string(Data))
	fmt.Println("Количество слов:", len(words))

	// Считаем среднее кол-во символов
	AllSymbols := 0
	for _, i := range words {
		// rune переводит число в unicode
		AllSymbols += len([]rune(i))
	}

	AVG := float64(AllSymbols) / float64(len(words))
	fmt.Println("Количество слов:", AVG)
}
