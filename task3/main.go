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

//	TASK 3
/*
Работа нескольких воркеров
Реализовать постоянную запись данных в канал (в главной горутине).

Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.

Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.
*/

func worker(wg sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}

//func main() {
//	var cg int
//	fmt.Scanln(&cg)
//
//	wg := sync.WaitGroup{}
//	ch := make(chan int)
//
//	for i := 0; i < cg; i++ {
//		wg.Add(1)
//		go worker(wg, ch)
//	}
//
//	ctx, cancel := context.WithCancel(context.Background())
//	go func() {
//		count := 0
//		for {
//			select {
//			case <-ctx.Done():
//				return
//			default:
//				count++
//				ch <- count
//				time.Sleep(300 * time.Millisecond)
//			}
//		}
//	}()
//
//	sigCh := make(chan os.Signal, 1)
//	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
//	<-sigCh
//
//	cancel()
//	wg.Wait()
//	close(ch)
//}

// если прям следовать заданию, то по условиям, ЗАПИСЬ осуществляется в главной горутине,
// ну это если я правильно понял, хотя мне вариант сверху больше нравиться

func main() {
	var cg int
	fmt.Scanln(&cg)

	wg := sync.WaitGroup{}
	ch := make(chan int)

	for i := 0; i < cg; i++ {
		wg.Add(1)
		go worker(wg, ch)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		<-sigCh
		cancel()
	}()

	count := 0
	for {
		select {
		case <-ctx.Done():
			close(ch)
			wg.Wait()
			return
		default:
			count++
			ch <- count
			time.Sleep(300 * time.Millisecond)
		}
	}
}
