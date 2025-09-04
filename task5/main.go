package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// TASK 5
/*(
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.

Подсказка: используйте time.After или таймер для ограничения времени работы.
*/

func main() {
	var timeout uint
	ch := make(chan int)

	fmt.Scanln(&timeout)

	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(timeout) * time.Second)
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					return
				}
				fmt.Println(v)
			case <-ctx.Done():
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		count := 0
		for {
			select {
			case ch <- count:
				count++
				time.Sleep(400 * time.Millisecond)
			case <-ctx.Done():
				//	закрываем там где пишем
				close(ch)
				return
			}
		}
	}()

	<-ctx.Done()
	//	И без этого будет работать, но на всякий случий лучше дождаться всех
	wg.Wait()
}
