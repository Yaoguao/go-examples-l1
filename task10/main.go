package main

import (
	"fmt"
	"math"
)

//	TASK 10
/*
Группировка температур
Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить эти значения в группы с шагом 10 градусов.

Пример: -20:{-25.4, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20:{24.5}, 30:{32.5}.

Пояснение: диапазон -20 включает значения от -20 до -29.9, диапазон 10 – от 10 до 19.9, и т.д.
Порядок в подмножествах не важен.
*/

func groupTemp(temps []float64) map[int][]float64 {
	res := make(map[int][]float64)
	var key int
	for _, temp := range temps {
		if temp <= 0 {
			key = int(10 * math.Ceil(temp/10))
		} else {
			key = int(10 * math.Floor(temp/10))
		}
		sl := res[key]
		res[key] = append(sl, temp)
	}
	return res
}

func main() {
	sl := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(groupTemp(sl))
}
