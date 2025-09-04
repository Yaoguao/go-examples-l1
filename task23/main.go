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

//	V2
//
// Можно и так
func RemoveAtIndexV2(s []int, index int) []int {
	if index < 0 || len(s) <= index {
		return s
	}
	return append(s[:index], s[index+1:]...)
}

// Написал 2 функции
func main() {
	sl := []int{1, 2, 3, 4}

	fmt.Println(sl) // [1, 2, 3, 4]

	// В таком случае лучше как и с append возвращать тому слайсу в котором удаляем
	tmp := RemoveAtIndexV2(sl, 3)

	tmp[0] = 33

	fmt.Println(sl) // [33, 2, 3, 4]
}
