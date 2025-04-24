package main

import (
	"fmt"
	"math"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func MaxABS(a, b []int) int {
	max := 0
	for i := range a {
		diff := int(math.Abs(float64(a[i] - b[i])))
		if diff > max {
			max = diff
		}
	}
	return max
}

func AvgABS(a, b []int) float64 {
	sum := 0
	for i := range a {
		sum += int(math.Abs(float64(a[i] - b[i])))
	}
	return float64(sum) / float64(len(a))
}

func ABSotl(a, b []int) []int {
	result := make([]int, len(a))
	for i := range a {
		result[i] = int(math.Abs(float64(a[i] - b[i])))
	}
	return result
}

func Random(Range, LowBorder, HighBorder int) []int {
	Mass := make([]int, Range)
	for i := 0; i < Range; i++ {
		Mass[i] = rand.Intn(HighBorder-LowBorder+1) + LowBorder
	}
	return Mass
}

func main() {
	// Ввод пользователя
	var Range, LowBorder, HighBorder int

	fmt.Print("Длина списка: ")
	fmt.Scan(&Range)
	fmt.Print("Нижняя граница генерации: ")
	fmt.Scan(&LowBorder)
	fmt.Print("Верхняя граница генерации: ")
	fmt.Scan(&HighBorder)

	// Вывод рандомных списков и отклонений
	ListOne := Random(Range, LowBorder, HighBorder)
	ListTwo := Random(Range, LowBorder, HighBorder)
	ListThree := Random(Range, LowBorder, HighBorder)
	ListFour := Random(Range, LowBorder, HighBorder)

	fmt.Printf("Список 1: %v\nСписок 2: %v\nСписок 3: %v\nСписок 4: %v\n", ListOne, ListTwo, ListThree, ListFour)

	fmt.Printf("\nМакс. отклонение (1 и 2): %d\n", MaxABS(ListOne, ListTwo))
	fmt.Printf("Среднее отклонение (1 и 2): %.2f\n", AvgABS(ListOne, ListTwo))
	fmt.Printf("Макс. отклонение (1 и 3): %d\n", MaxABS(ListOne, ListThree))
	fmt.Printf("Среднее отклонение (1 и 3): %.2f\n", AvgABS(ListOne, ListThree))
	fmt.Printf("Макс. отклонение (1 и 4): %d\n", MaxABS(ListOne, ListFour))
	fmt.Printf("Среднее отклонение (1 и 4): %.2f\n", AvgABS(ListOne, ListFour))

	// Задаем нужные графики
	abs12 := ABSotl(ListOne, ListTwo)
	abs13 := ABSotl(ListOne, ListThree)
	abs14 := ABSotl(ListOne, ListFour)

	// Пишем названия для графика
	p := plot.New()
	p.Title.Text = "Абсолютные отклонения"
	p.X.Label.Text = "Индекс элемента"
	p.Y.Label.Text = "Значение отклонения"

	// Создание трех срезов (X, Y) для отрезков
	pts12 := make(plotter.XYs, len(abs12))
	pts13 := make(plotter.XYs, len(abs13))
	pts14 := make(plotter.XYs, len(abs14))
	// Этот цикл заполняет строит график (Заполняет точками)
	for i := 0; i < Range; i++ {
		pts12[i].X = float64(i)
		pts12[i].Y = float64(abs12[i])
		pts13[i].X = float64(i)
		pts13[i].Y = float64(abs13[i])
		pts14[i].X = float64(i)
		pts14[i].Y = float64(abs14[i])
	}

	// Задает цвета для графиков
	line12, _ := plotter.NewLine(pts12)
	line12.Color = plotutil.Color(0)
	line12.LineStyle.Width = vg.Points(2)

	line13, _ := plotter.NewLine(pts13)
	line13.Color = plotutil.Color(1)
	line13.LineStyle.Width = vg.Points(2)

	line14, _ := plotter.NewLine(pts14)
	line14.Color = plotutil.Color(2)
	line14.LineStyle.Width = vg.Points(2)

	// Пишем текст, какие линии, за что отвечают
	p.Add(line12, line13, line14)
	p.Legend.Add("Отклонение 1 и 2", line12)
	p.Legend.Add("Отклонение 1 и 3", line13)
	p.Legend.Add("Отклонение 1 и 4", line14)

	// Сохранения файла
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "abs_diff_plot.png"); err != nil {
		panic(err)
	}
}
