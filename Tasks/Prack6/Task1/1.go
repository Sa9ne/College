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
	// Смотрим длину массива и вычетам один т.к. массив начинается в 0
	for i := 0; i < MasRange-1; i++ {
		// Делаем второй вложенный цикл
		for j := 0; j < MasRange-i-1; j++ {
			if Mas[j] > Mas[j+1] {
				// Меняем элементы местами
				Mas[j], Mas[j+1] = Mas[j+1], Mas[j]
			}
		}
	}
	// Хочу глянуть сам массив
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
