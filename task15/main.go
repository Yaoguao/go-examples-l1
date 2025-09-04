package main

import (
	"fmt"
	"strings"
)

//	TASK 15

/*
Небольшой фрагмент кода — проблемы и решение
Рассмотреть следующий код и ответить на вопросы: к каким негативным последствиям он может привести и как это исправить?

Приведите корректный пример реализации.

var justString string

func someFunc() {
  v := createHugeString(1 &lt;&lt; 10) * я не понимаю что такое 1 &lt;&lt; 10
  justString = v[:100]
}

func main() {
  someFunc()
}
Вопрос: что происходит с переменной justString?
*/

var justString string

func someFunc() {
	v := createHugeString(1024 * 1024)
	// данные ссылются на одну область пямяти, может быть не ожиданной проблемой
	// так как строки это неизменяемы массив байт, мы извлекаем хоть и новую строку,
	// но дальнейший массив(на которы мы ссылаемся) никуда не деваеться и сборщик не сможет его очистить(так как он используеться)
	// и "емкость" justString будет колосально огромной, что может быть затратно
	// justString = v[:100]
	//решить проблему можно так
	// justString = string([]byte(v[:100]))
	// или так
	justString = strings.Clone(v[:100])

	fmt.Println(justString)
}

func createHugeString(len int) string {
	return string(make([]byte, len))
}

func main() {
	someFunc()
}
