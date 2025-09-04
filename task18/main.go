package main

import (
	"fmt"
	"sync"
)

//	TASK 18

/*
Конкурентный счетчик
Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
По завершению программы структура должна выводить итоговое значение счётчика.

Подсказка: вам понадобится механизм синхронизации, например, sync.Mutex или sync/Atomic для безопасного инкремента.
*/

type Counter struct {
	mx    sync.RWMutex
	Count int
}

func (c *Counter) Increment() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.Count++
}
func (c *Counter) Decrement() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.Count--
}

func main() {
	wg := sync.WaitGroup{}
	c := Counter{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}
	wg.Wait()

	fmt.Println(c.Count)
}
