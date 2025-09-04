package main

import (
	"fmt"
	"time"
)

//	TASK25

/*
Своя функция Sleep
Реализовать собственную функцию sleep(duration) аналогично встроенной функции time.Sleep,
которая приостанавливает выполнение текущей горутины.

Важно: в отличии от настоящей time.Sleep,
ваша функция должна именно блокировать выполнение (например, через таймер или цикл),
а не просто вызывать time.Sleep :) — это упражнение.

Можно использовать канал + горутину, или цикл на проверку времени (не лучший способ, но для обучения).
*/

func MySleep(d time.Duration) {
	timer := time.NewTimer(d)
	<-timer.C
	// time.After возвращает канал
	// <-time.After(d)
}

// не лучший способ :D
//func MySleep(d time.Duration) {
//	now := time.Now()
//	for time.Since(now) < d {
//	}
//}

func main() {
	go func() {
		fmt.Println("WOW")
		MySleep(1 * time.Second)
		fmt.Println("WOW")
	}()

	MySleep(2 * time.Second)
}
