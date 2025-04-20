package main

import (
	"fmt"
	"math/rand"
)

func InfoFIO(Students [4]string, Math []int, Russ []int, Phys []int) {
	// Спрашиваем чьи оценки вывести
	var ChooseStudent int

	fmt.Println("\nОценку какого студента вывести?")
	for i := 0; i < len(Students); i++ {
		fmt.Printf("%d - %v\n", i+1, Students[i])
	}
	fmt.Scan(&ChooseStudent)
	fmt.Print("\n")

	// Выводим данные
	if ChooseStudent > len(Students) {
		fmt.Println("Такого студента не существует")
		return
	}

	fmt.Println(Students[ChooseStudent-1])
	fmt.Println("Математика:", Math[ChooseStudent-1])
	fmt.Println("Русский:", Russ[ChooseStudent-1])
	fmt.Println("Физ-ра:", Phys[ChooseStudent-1])

}

func InfoSubject(Students [4]string, Math []int, Russ []int, Phys []int) {
	// Вывод пользователю
	var ChooseSubject, ChooseAssessment int

	fmt.Println("\nПо какому предмету?")
	fmt.Println("1 - Математика")
	fmt.Println("2 - Русский")
	fmt.Print("3 - Физ-ра\n")
	fmt.Scan(&ChooseSubject)

	fmt.Print("\nКакую оценку? (от 2 до 5)\n\n")
	fmt.Scan(&ChooseAssessment)

	// Проверка ввода
	if ChooseAssessment < 2 || ChooseAssessment > 5 {
		fmt.Println("Не верно введенная оценка")
		return
	}

	// Вывод текста для различной ситуации
	if ChooseSubject == 1 {
		fmt.Printf("Оценка %d по математике: \n", ChooseAssessment)
	} else if ChooseSubject == 2 {
		fmt.Printf("Оценка %d по русскому: ", ChooseAssessment)
	} else if ChooseSubject == 3 {
		fmt.Printf("Оценка %d по физ-ре: ", ChooseAssessment)
	} else {
		fmt.Println("Не верно введенный предмет")
		return
	}

	// Нахождение нужных оценок
	for i := 0; i < len(Students); i++ {
		if ChooseSubject == 1 {
			if Math[i] == ChooseAssessment {
				fmt.Println(Students[i])
			}
		} else if ChooseSubject == 2 {
			if Russ[i] == ChooseAssessment {
				fmt.Println(Students[i])
			}
		} else if ChooseSubject == 3 {
			if Phys[i] == ChooseAssessment {
				fmt.Println(Students[i])
			}
		} else {
			fmt.Println("Не верно введенный предмет")
			return
		}
	}

}

func main() {
	// Задали список студентов (Надеюсь я правильно понял задание)
	Students := [4]string{"Коротков", "Третьяков", "Петрова", "Воробьева"}

	// Рандомизируем оценки
	var Math, Russ, Phys []int
	var Random int

	// Мне кажется можно оптимизировать как то лучше
	for i := 0; i < 4; i++ {
		Random = rand.Intn(6-2) + 2
		Math = append(Math, Random)
	}

	for i := 0; i < 4; i++ {
		Random = rand.Intn(6-2) + 2
		Russ = append(Russ, Random)
	}

	for i := 0; i < 4; i++ {
		Random = rand.Intn(6-2) + 2
		Phys = append(Phys, Random)
	}

	// Вывод пользователю
	fmt.Println("Список студентов:", Students)
	fmt.Println("Список сгенерированных оценок: ")
	fmt.Println("Математика:", Math)
	fmt.Println("Русский:", Russ)
	fmt.Println("Физкультура:", Phys)

	// Текст для выбора задачи
	var ChooseTask int
	Stop := false

	for !Stop {
		fmt.Println("\nВыберите задачу:")
		fmt.Println("1 - Вывести оценки по фамилии")
		fmt.Println("2 - Вывести студентов по оценке")
		fmt.Println("0 - Выход")
		fmt.Scan(&ChooseTask)

		// Выбор задачки
		if ChooseTask == 1 {
			InfoFIO(Students, Math, Russ, Phys)
		} else if ChooseTask == 2 {
			InfoSubject(Students, Math, Russ, Phys)
		} else if ChooseTask == 0 {
			Stop = true
		} else {
			fmt.Println("Вы ввели не существующую задачу")
		}
	}
}
