package main

import "fmt"

//	TASK 11
/*
Пересечение множеств
Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) — т.е. вывести элементы,
присутствующие и в первом, и во втором.

Пример:
A = {1,2,3}
B = {2,3,4}
Пересечение = {2,3}
*/

func intersection(nums1 []int, nums2 []int) []int {
	var hmap = make(map[int]int)
	var res []int

	for _, num := range nums1 {
		if hmap[num] > 0 {
			continue
		} else {
			hmap[num] = 1
		}
	}

	for _, num := range nums2 {
		if hmap[num] > 0 {
			hmap[num] += 1
		} else {
			continue
		}
	}

	for num, count := range hmap {
		if count > 1 {
			res = append(res, num)
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 2, 3}

	nums2 := []int{2, 3, 4}

	fmt.Println(intersection(nums1, nums2))
}
