package main

import (
	"fmt"
)

// Пишем структуру Stack
type Stack struct {
	// rune что бы принимать в unicode
	Data []rune
}

// Добавляем символ в стек
func (s *Stack) Push(r rune) {
	s.Data = append(s.Data, r)
}

// Добавим проверку на пустой стек
func (s *Stack) IsEmpty() bool {
	return len(s.Data) == 0
}

// Достаем верхнюю скобку
func (s *Stack) Pop() rune {
	// Проверка на пустой стек
	if s.IsEmpty() {
		return 0
	}

	// Берём последний элемент
	last := s.Data[len(s.Data)-1]

	// Удаляем его из среза
	s.Data = s.Data[:len(s.Data)-1]

	return last
}

// Проверка скобочной последовательности
func IsValidBrackets(Message string) bool {
	// Создаем пустой стек
	stack := Stack{}

	// Создаем хеш мапу для соответствия скобок
	Skobki := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// Проходимся по всей длине message
	for _, char := range Message {
		// Если скобка открывающая, то добавляем символ в стек
		if char == '(' || char == '[' || char == '{' {
			stack.Push(char)
		} else if char == ')' || char == ']' || char == '}' {
			// Если стек не пустой, мы сравниваем скобки из хеш мапы и если они не совпадают, то это ошибка
			if stack.IsEmpty() || stack.Pop() != Skobki[char] {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

func main() {
	var Message string
	var Stop bool

	// Пока работа true
	for !Stop {
		fmt.Print("Введите последовательность символов: ")
		fmt.Scan(&Message)

		if Message == "стоп" {
			Stop = true
		} else if IsValidBrackets(Message) {
			fmt.Println("Последовательность корректна")
		} else {
			fmt.Println("Последовательность НЕ корректна")
		}
	}
}
