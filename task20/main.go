package main

import (
	"fmt"
	"strings"
)

//	TASK 20

/*
Разворот слов в предложении
Разработать программу, которая переворачивает порядок слов в строке.

Пример: входная строка:

«snow dog sun», выход: «sun dog snow».

Считайте, что слова разделяются одиночным пробелом.
Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».
* Честно говоря я не знаю как сделать именно на месте, без доп слайса. Это разве возможно?
*/

// Решение похожее на прошлую задачу
func ReverseWords(str string) string {
	// Либо strings.Split()
	sl := strings.Fields(str)
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	return strings.Join(sl, " ")
}

// Решение с билдером
//func ReverseWords(str string) string {
//	var builder strings.Builder
//	sl := strings.Fields(str)
//	for i := len(sl) - 1; i >= 0; i-- {
//		builder.WriteString(sl[i])
//		builder.WriteString(" ")
//	}
//	return strings.TrimSpace(builder.String())
//}

func main() {
	str := "snow dog sun"
	fmt.Println(ReverseWords(str))
}
