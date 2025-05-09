package main

import (
	"fmt"
	"os"
	"unicode"
)

func caesarEncrypt(Data string, Shift int) string {
	var result []rune

	for _, char := range Data {
		if unicode.IsLetter(char) {
			var base rune
			// по умолчанию — кириллица
			alphabetSize := 32

			// Определим базу для латиницы
			if unicode.Is(unicode.Latin, char) {
				alphabetSize = 26
				if unicode.IsUpper(char) {
					base = 'A'
				} else {
					base = 'a'
				}
			} else if unicode.Is(unicode.Cyrillic, char) {
				if unicode.IsUpper(char) {
					base = 'А'
				} else {
					base = 'а'
				}
			} else {
				// не буква, пропускаем
				result = append(result, char)
				continue
			}

			// Смещаем букву
			shifted := (int(char-base) + Shift) % alphabetSize
			if shifted < 0 {
				shifted += alphabetSize
			}
			result = append(result, rune(int(base)+shifted))
		} else {
			// Прочие символы остаются без изменений
			result = append(result, char)
		}
	}

	return string(result)
}

func caesarDecrypt(Data string, Shift int) string {
	return caesarEncrypt(Data, -Shift)
}

func main() {
	// Ввод пользователя
	Shift := 0
	fmt.Print("Введите сдвиг: ")
	fmt.Scan(&Shift)

	// Чтение файла
	Data, err := os.ReadFile("text.txt")
	if err != nil {
		panic(fmt.Sprintln("При чтении файла нашлась ошибка: ", err))
	}

	work := true
	for work {
		Choice := 0
		fmt.Print("Выберите режим (1 - шифровать, 2 - дешифровать): ")
		fmt.Scan(&Choice)

		if Choice == 1 {
			// Вызов функции зашифровки
			encrypted := caesarEncrypt(string(Data), Shift)

			err = os.WriteFile("output.txt", []byte(encrypted), 0644)
			if err != nil {
				panic(fmt.Sprintln("Ошибка при записи в файл:", err))
			}
		} else if Choice == 2 {
			// Чтение файла
			Data, err := os.ReadFile("output.txt")
			if err != nil {
				panic(fmt.Sprintln("При чтении файла нашлась ошибка: ", err))
			}

			// Вызов функции дешифровки
			encrypted := caesarDecrypt(string(Data), Shift)

			err = os.WriteFile("output.txt", []byte(encrypted), 0644)
			if err != nil {
				panic(fmt.Sprintln("Ошибка при записи в файл:", err))
			}
		} else {
			work = false
		}
	}
}
