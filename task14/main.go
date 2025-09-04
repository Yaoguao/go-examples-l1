package main

import (
	"fmt"
)

//	TASK 14
/*
Определение типа переменной в runtime
Разработать программу, которая в runtime способна определить тип переменной,
переданной в неё (на вход подаётся interface{}). Типы, которые нужно распознавать: int, string, bool, chan (канал).

Подсказка: оператор типа switch v.(type) поможет в решении.
*/

func determinant(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan string: // ????
		fmt.Println("chan")
	default:
		fmt.Println("I dont know")
	}
}

// Вариант с рефлексией, тогда можно определить любой канал
//func determinant(i interface{}) {
//	switch v := reflect.ValueOf(i); v.Kind() {
//	case reflect.Int:
//		fmt.Println("int")
//	case reflect.String:
//		fmt.Println("string")
//	case reflect.Bool:
//		fmt.Println("bool")
//	case reflect.Chan:
//		fmt.Println("chan")
//	default:
//		fmt.Println("I dont know")
//	}
//}

func main() {
	ch := make(chan string)
	determinant(ch)
}
