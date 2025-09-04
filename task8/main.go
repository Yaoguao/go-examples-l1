package main

import "fmt"

//	TASK 8

/*
Установка бита в числе
Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.

Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).

Подсказка: используйте битовые операции (|, &^).
*/

func ConvertIntToBits(n int) []uint8 {
	bits := make([]uint8, 8)
	for i := 7; i >= 0; i-- {
		bits[7-i] = uint8((n >> i) & 1)
	}
	return bits
}

func ConvertBitsToInt(bits []uint8) int {
	if len(bits) != 8 {
		panic("не равно 8 битам")
	}
	length := 0
	for i := 0; i < 8; i++ {
		length = (length << 1) | int(bits[i])
	}
	return length
}

// ну вроде правильно
func main() {
	var res, index int
	// ограничусь 8 битами
	fmt.Println("ведите число(до 255):")
	_, err := fmt.Scanln(&res)
	if err != nil {
		return
	}

	sb := ConvertIntToBits(res)
	fmt.Printf("число в двоичном виде: %v\n", sb) // [0 0 0 0 0 1 0 0]
	fmt.Println("индекс от 1 до 8 с право на лево")
	_, err = fmt.Scanln(&index)
	if err != nil {
		return
	}

	if index < 0 || index > 8 {
		fmt.Println("неверно")
		return
	}

	revIdx := len(sb) - index

	if sb[revIdx] != 0 {
		sb[revIdx] = 0
	} else {
		sb[revIdx] = 1
	}
	fmt.Printf("число в двоичном виде: %v\n", sb)
	res = ConvertBitsToInt(sb)
	fmt.Printf("число: %d\n", res)
	fmt.Println("ПОКА!!")
}
