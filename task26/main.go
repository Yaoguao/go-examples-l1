package main

import (
	"fmt"
	"strings"
)

//	TASK 26

/*
Уникальные символы в строке
Разработать программу, которая проверяет,
что все символы в строке встречаются один раз(т.е. строка состоит из уникальных символов).

Вывод: true, если все символы уникальны, false, если есть повторения.
Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.

Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.

Подумайте, какой структурой данных удобно воспользоваться для проверки условия.
*/
func Unique(s string) bool {
	ch := make(map[rune]struct{})
	for _, val := range strings.ToLower(s) {
		if _, ok := ch[val]; ok {
			return false
		}
		ch[val] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(Unique("acfgA"))
}
