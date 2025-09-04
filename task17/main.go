package main

import "fmt"

//	TASK 17

/*
Бинарный поиск
Реализовать алгоритм бинарного поиска встроенными методами языка.
Функция должна принимать отсортированный слайс и искомый элемент, возвращать индекс элемента или -1, если элемент не найден.

Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.
*/
func BinarySearch(val int, sl []int) int {
	left, right, mid := 0, len(sl)-1, 0
	for left <= right {
		mid = left + (right-left)/2
		if sl[mid] == val {
			return mid
		} else if sl[mid] < val {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(BinarySearch(3, sl))
}
