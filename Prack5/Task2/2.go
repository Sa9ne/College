package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	Data, err := os.ReadFile("text.txt")
	if err != nil {
		panic(fmt.Sprintln("При чтении файла нашлась ошибка: ", err))
	}

	// Создали через make, так как это ссылочный тип данных (ссылается на место в памяти, а не создает переменную)
	// map используется для создания хеш-мапы
	LetterStats := make(map[rune]int)

	for _, char := range string(Data) {
		// Проверка на число
		if unicode.IsLetter(char) {
			// Приводим к нижнему регистру
			char = unicode.ToLower(char)
			LetterStats[char]++
		}
	}

	for letter, count := range LetterStats {
		fmt.Printf("%c: %d\n", letter, count)
	}
}
