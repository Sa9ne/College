package main

import (
	"fmt"
	"strings"
)

// Создаем структуру Character в которой будут характеристики сущности
type Character struct {
	X int
	Y int
}

// Создаем нового персонажа
func NewCharacter(StartX int, StartY int) Character {
	return Character{X: StartX, Y: StartY}
}

// Рисуем игровое поле
func (ctx *Character) PrintBoard() {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if ctx.X == x && ctx.Y == y {
				fmt.Print("🔵")
			} else {
				fmt.Print("⬜")
			}
		}
		fmt.Println()
	}
}

func (ctx *Character) MoveUp() {
	if ctx.Y > 0 {
		ctx.Y--
	} else {
		fmt.Println("Невозможно двигаться вверх")
	}
}

func (ctx *Character) MoveDown() {
	if ctx.Y < 7 {
		ctx.Y++
	} else {
		fmt.Println("Невозможно двигаться вниз")
	}
}

func (ctx *Character) MoveLeft() {
	if ctx.X > 0 {
		ctx.X--
	} else {
		fmt.Println("Невозможно двигаться влево")
	}
}
func (ctx *Character) MoveRight() {
	if ctx.X < 7 {
		ctx.X++
	} else {
		fmt.Println("Невозможно двигаться вправо")
	}
}

func main() {
	fmt.Println("Функционал приложения:\n\tВверх\n\tВниз\n\tВправо\n\tВлево\n\tСтоп")

	var X, Y int
	fmt.Print("Введите стартовую позицию по X: ")
	fmt.Scan(&X)
	fmt.Print("Введите стартовую позицию по X: ")
	fmt.Scan(&Y)

	char := NewCharacter(X, Y)

	if X < 0 || X >= 8 || Y < 0 || Y >= 8 {
		fmt.Println("Поле не настолько большое")
		return
	}

	var Command string

	char.PrintBoard()

	for {
		fmt.Print("\nВведите команду: ")
		fmt.Scan(&Command)

		Commands := strings.ToLower(Command)

		if Commands == "вверх" {
			char.MoveUp()
		} else if Commands == "вниз" {
			char.MoveDown()
		} else if Commands == "вправо" {
			char.MoveRight()
		} else if Commands == "влево" {
			char.MoveLeft()
		} else if Commands == "стоп" {
			fmt.Println("Конец данной игры")
			return
		}

		char.PrintBoard()
	}
}
