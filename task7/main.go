package main

import (
	"fmt"
	"sync"
)

//	TASK 7

/*
Конкурентная запись в map
Реализовать безопасную для конкуренции запись данных в структуру map.

Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).

Проверьте работу кода на гонки (util go run -race).
*/
type MapCon struct {
	mx *sync.RWMutex
	h  map[string]int
}

func (ms *MapCon) Set(key string, value int) {
	ms.mx.Lock()
	defer ms.mx.Unlock()
	ms.h[key] = value
}

func (ms *MapCon) Get(key string) (int, bool) {
	ms.mx.RLock()
	val, ok := ms.h[key]
	ms.mx.RUnlock()
	return val, ok
}

func main() {
	wg := sync.WaitGroup{}

	m := MapCon{
		mx: &sync.RWMutex{},
		h:  make(map[string]int),
	}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			v, _ := m.Get("TEST")
			m.Set("TEST", 1+v)
		}(i)
	}

	wg.Wait()
	fmt.Println(m.h)
}
