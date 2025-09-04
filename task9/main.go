package main

import (
	"fmt"
)

//	TASK 9
/*
Конвейер чисел
Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2.
После этого данные из второго канала должны выводиться в stdout.
То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка.
Убедитесь, что чтение из второго канала корректно завершается.
*/

func writer(arr []int) <-chan int {
	ch := make(chan int)

	go func() {
		for val := range arr {
			ch <- val
		}
		close(ch)
	}()

	return ch
}

func doubler(inputCh <-chan int) <-chan int {
	outCh := make(chan int)

	go func() {
		for val := range inputCh {
			outCh <- val * 2
		}
		close(outCh)
	}()

	return outCh
}

func reader(inputCh <-chan int) {
	for val := range inputCh {
		fmt.Println(val)
	}
}

func main() {
	arr := []int{1, 3, 5, 7, 8, 9, 5, 4, 3, 2, 76, 34, 3}
	reader(doubler(writer(arr)))
}
