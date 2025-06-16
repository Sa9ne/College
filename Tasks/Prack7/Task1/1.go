package main

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func Sort(Mas []int, Low int, High int) {
	// В go алгоритм быстрой сортировки используется по умолчанию в стандартной библиотеке))
	// sort.Ints(Mas)

	// Напишу ручками

	// Для остановки рекурсии
	if Low >= High {
		return
	}

	fmt.Println(Mas)

	Wall := Low        // Стена
	Pivot := Mas[High] // Последний элемент массива

	for i := Low; i < High; i++ {
		if Mas[i] <= Pivot {
			// Меняем местами значения
			Mas[Wall], Mas[i] = Mas[i], Mas[Wall]

			// Переносим стену дальше, что бы первый элемент массива больше не трогали
			Wall++

			// Для наглядности выведу сам массив
			fmt.Println(Mas)
		}
	}

	// Переноси опорный элемент (Pivot) в начало нашего массива
	Mas[Wall], Mas[High] = Mas[High], Mas[Wall]

	// Вызываем рекурсию, эта рекурсия сортирует правую и левую часть
	Sort(Mas, Low, Wall-1)
	Sort(Mas, Wall+1, High)
}

func main() {
	var Lower, Upper int
	var points plotter.XYs
	fmt.Print("Введите нижнюю границу: ")
	fmt.Scan(&Lower)

	fmt.Print("Введите верхнюю границу: ")
	fmt.Scan(&Upper)

	for i := 0; i < 3; i++ {
		// Генерация рандома и занесения в массив
		var Mas []int

		// Взял 6 как длину массива, не очень понял что именно нужно взять
		for i := 0; i < 6; i++ {
			Random := rand.Intn(Upper-Lower+1) + Lower
			Mas = append(Mas, Random)
		}

		// Засекаем время
		TimeStart := time.Now()
		Sort(Mas, 0, len(Mas)-1)
		Duration := time.Since(TimeStart).Seconds()

		points = append(points, plotter.XY{
			X: float64(i + 1), // номер попытки
			Y: Duration,
		})

		fmt.Printf("Прошло времени с момента начала: %.6f секунд\n", Duration)
	}

	// Рисуем график
	p := plot.New()
	p.Title.Text = "Время сортировки"
	p.X.Label.Text = "Попытка"
	p.Y.Label.Text = "Время (секунды)"

	err := plotutil.AddLinePoints(p, points)
	if err != nil {
		panic(err)
	}

	// Сохраняем график в файл
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "output.png"); err != nil {
		panic(err)
	}

	fmt.Println("График сохранён в output.png")
}
