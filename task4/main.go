package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//	TASK 4
/*
Завершение по Ctrl+C
Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).

Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.

Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении.
*/

func worker(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	//	Так как цикл for range работает так, что итерация по каналу происходит до момента закрытия, это и будет моментом его завершения(закрытие)
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	var cg int
	fmt.Scanln(&cg)

	//	Создаем группу для отслеживания горутин
	wg := sync.WaitGroup{}
	ch := make(chan int)

	for i := 0; i < cg; i++ {
		wg.Add(1)
		go worker(&wg, ch)
	}

	//	Взял решение с прошлой задачи
	//	просто shutdown
	//	Использую контекст для отслеживания о прекращении работы
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigCh := make(chan os.Signal, 1)
		//	Новосозданный канал передаем в функцию notify
		signal.Notify(sigCh, syscall.SIGINT)

		//	Стоим и ждем пока нам не скажут, когда поступил сигнал
		<-sigCh
		//	Завершаем работу контекста
		cancel()
	}()

	count := 0
	for {
		select {
		//	А тут ловим)
		case <-ctx.Done():
			close(ch)
			//	Дожидаемся когда все наши воркеры завершаться
			wg.Wait()
			return
		default:
			count++
			ch <- count
			time.Sleep(300 * time.Millisecond)
		}
	}
}
