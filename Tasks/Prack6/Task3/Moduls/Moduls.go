package moduls

import (
	"fmt"
	"math/rand"
	"sort"
)

func Random(Upper int, Lower int, Posled int) []int {
	// Определение слайса в котором значения []int, изначально находиться 0 элементов, емкость слайса
	Mas := make([]int, 0, Posled)

	// Присвоение рандома
	for i := 0; i < Posled; i++ {
		Random := rand.Intn(Upper-Lower+1) + Lower
		Mas = append(Mas, Random)
	}

	// Сортировка массива
	sort.Ints(Mas)

	return Mas
}

func BinarySearch(Mas []int, FoundNum int) int {

	Low := 0
	High := len(Mas) - 1
	Iteration := 1

	fmt.Println("\nИтерация\tLow\tMid\tHigh\tЗначение")
	for Low <= High {
		Mid := (Low + High) / 2
		fmt.Printf("%d\t\t%d\t%d\t%d\t%d\n", Iteration, Low, Mid, High, Mas[Mid])

		// Если центр окажется числом, то мы выводим индекс
		if Mas[Mid] == FoundNum {
			return Mid
		} else if Mas[Mid] < FoundNum { // Если центр массива меньше заданного числа, то его там точно нет
			Low = Mid + 1
		} else if Mas[Mid] > FoundNum { // Если центр массива больше числа, то правее его точно не будет
			High = Mid - 1
		}
		// -1 и +1 нужны для того что бы убрать число, которое мы уже явно проверили (Mas[Mid] == FoundNum)
		Iteration++
	}
	// Если число не будет найдено, возвращается -1 как про
	return -1
}
