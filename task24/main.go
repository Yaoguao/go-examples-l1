package main

import (
	"fmt"
	"math"
)

//	TASK24
/*
Расстояние между точками
Разработать программу нахождения расстояния между двумя точками на плоскости.
Точки представлены в виде структуры Point с инкапсулированными (приватными) полями x, y (типа float64) и конструктором.
Расстояние рассчитывается по формуле между координатами двух точек.

Подсказка: используйте функцию-конструктор NewPoint(x, y), Point и метод Distance(other Point) float64.
*/
type Point struct {
	x, y int
}

func New(x, y int) Point {
	return Point{x, y}
}

func (p Point) Distance(other Point) float64 {
	dx := float64(other.x - p.x)
	dy := float64(other.y - p.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p := New(0, 0)
	p2 := New(0, 4)
	fmt.Println(p.Distance(p2))
}
