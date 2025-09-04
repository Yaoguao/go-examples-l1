package main

//	TASK 6
/*
Реализовать все возможные способы остановки выполнения горутины.

Классические подходы: выход по условию, через канал уведомления, через контекст, прекращение работы runtime.Goexit() и др.

Продемонстрируйте каждый способ в отдельном фрагменте кода.
*/

// Выход по условию
/*
func main() {
	wg := sync.WaitGroup{}
	flag := false
	wg.Add(1)
	go func() {
		defer wg.Done()
		for !flag {
			fmt.Println("go running")
			time.Sleep(200 * time.Millisecond)
		}
	}()

	time.Sleep(400 * time.Millisecond)
	flag = true
	wg.Wait()

}
*/

// Выход через канал

/*
func main() {
	wg := sync.WaitGroup{}
	ch := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch:
				return
			default:
				fmt.Println("go running")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(400 * time.Millisecond)
	// или так ch <- struct{}{}
	close(ch)
	wg.Wait()

}
*/

// Выход через контекст
/*
func main() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			//	Можно ли сказать что почти тоже самое что и через канал?)
			case <-ctx.Done():
				return
			default:
				fmt.Println("go running")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(400 * time.Millisecond)
	cancel()
	wg.Wait()

}
*/

// Выхот тоже через контекст но используя таймаут
/*
func main() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("go running")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(600 * time.Millisecond)
	wg.Wait()

}
*/

// Выхот runtime.Goexit()
/*
func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 4; i++ {
			if i == 3 {
				runtime.Goexit()
			}
			fmt.Println("go running")
			time.Sleep(200 * time.Millisecond)
		}
	}()

	wg.Wait()

}
*/

//	Выход panic
/*
func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if p := recover(); p != nil {
				fmt.Println("dont worry")
			}
		}()
		for i := 0; i < 4; i++ {
			if i == 3 {
				panic("paniic")
			}
			fmt.Println("go running")
			time.Sleep(200 * time.Millisecond)
		}
	}()

	wg.Wait()
}
*/

// Выход через сигнал
/*
func main() {
	wg := sync.WaitGroup{}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-sigch:
				return
			default:
				fmt.Println("go running")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}
*/
