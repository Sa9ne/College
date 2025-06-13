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

func Sort(Mas []int) {
	MasRange := len(Mas)
	for i := 1; i < MasRange; i++ {
		key := Mas[i]
		j := i - 1

		// Перемещаем элементы Mas[0..i-1], которые больше key, на одну позицию вперед
		for j >= 0 && Mas[j] > key {
			Mas[j+1] = Mas[j]
			j = j - 1
		}
		Mas[j+1] = key
	}
	fmt.Println(Mas)
}

func main() {
	var Lower, Upper int
	var points plotter.XYs
	fmt.Println("Введите нижнюю границу")
	fmt.Scan(&Lower)

	fmt.Println("Введите верхнюю границу")
	fmt.Scan(&Upper)

	for i := 0; i < 3; i++ {
		// Генерация рандома и занесения в массив
		var Mas []int

		// Взял 20 как длину массива, не очень понял что именно нужно взять
		for i := 0; i < 20; i++ {
			Random := rand.Intn(Upper-Lower+1) + Lower
			Mas = append(Mas, Random)
		}

		// Засекаем время
		TimeStart := time.Now()
		Sort(Mas)
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

	err := plotutil.AddLinePoints(p, "Сортировка", points)
	if err != nil {
		panic(err)
	}

	// Сохраняем график в файл
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "output.png"); err != nil {
		panic(err)
	}

	fmt.Println("График сохранён в output.png")
}
