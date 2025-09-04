package main

import "fmt"

//	TASK 23

/*
Удаление элемента слайса
Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.

Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента
(copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.

*/

//	V1
/*
func RemoveAtIndex(s []int, index int) []int {
	res := make([]int, 0, len(s))
	for idx, val := range s {
		if idx == index {
			continue
		}
		res = append(res, val)
	}
	return res
}
*/

// V2
func RemoveAtIndexV2(s []int, index int) []int {
	if index < 0 || len(s) <= index {
		return s
	}
	return append(s[:index], s[index+1:]...)
}

func main() {
	sl := []int{1, 2, 3, 4}

	fmt.Println(sl) // [1, 2, 3, 4]

	sl = RemoveAtIndexV2(sl, 1)

	fmt.Println(sl) // [1, 3, 4]
}
