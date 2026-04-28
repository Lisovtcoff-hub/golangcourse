package main

import (
	"fmt"
)

func main() {
	var names = []string{"Tom", "Alice", "Kate"}

	// Рисовать: Передача слайса в функцию
	printer := func(s []string) {
		fmt.Println(s)
	}

	// Структура слайса копируется!
	printer(names)

	// Рисовать: Возврат из функции
	returner := func() []string {
		s := []string{"Tom", "Alice", "Kate"}
		return s
	}

	// Получили копию структуры (копию указателя на массив и копии len и cap)
	s := returner()
	_ = s
}
