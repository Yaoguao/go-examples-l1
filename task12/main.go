package main

import "fmt"

//	TASK 12
/*
Собственное множество строк
Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.

Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

*/

func newSet(sl []string) map[string]struct{} {
	res := make(map[string]struct{}, len(sl))
	for _, val := range sl {
		res[val] = struct{}{}
	}
	return res
}

func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}
	set := newSet(sl)
	for key, _ := range set {
		fmt.Println(key)
	}
	fmt.Println(set)
}
