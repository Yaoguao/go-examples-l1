package main

import "fmt"

//	TASK 16
/*
Быстрая сортировка (quicksort)
Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.

Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел.
Для выбора опорного элемента можно взять середину или первый элемент.
*/

func QSort(sl []int) []int {
	res := make([]int, len(sl))
	copy(res, sl)
	quicksort(res, 0, len(res)-1)
	return res
}

func quicksort(sl []int, low, high int) {
	if low < high {
		p := func(sl []int, low, high int) int { // можно вынести в отдельную функцию
			pivot := sl[(low+high)/2]
			i := low
			j := high
			for {
				for sl[i] < pivot {
					i++
				}
				for sl[j] > pivot {
					j--
				}
				if i >= j {
					return j
				}
				sl[i], sl[j] = sl[j], sl[i]
				i++
				j--
			}
		}(sl, low, high)
		quicksort(sl, low, p)
		quicksort(sl, p+1, high)
	}
}

func main() {
	sl := []int{2, 5, 7, 9, 4, 10, 43, 5, 1, 788, 56, 43, 12, 8, 767, 4, 6, 7, 9, 5}
	sl = QSort(sl)
	fmt.Println(sl)
}
