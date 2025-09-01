package main

import "fmt"

//	TASK 1
/*
Дана структура Human (с произвольным набором полей и методов).

Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.
*/

type Human struct {
	Name string
	Age  int
}

func (h Human) HelloHuman() {
	fmt.Printf("Hello, %s\n", h.Name)
}

type Action struct {
	Human
}

func (a Action) ActionHuman() {
	a.HelloHuman()
}

func main() {
	h := Human{Name: "Ass"}
	a := Action{h}

	a.ActionHuman()
}
