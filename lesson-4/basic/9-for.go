package main

import (
	"fmt"
	"time"
)

// Циклы
func main() {
	// for - это единственный цикл в Go

	{ // Бесконечный цикл с условием выхода
		var counter int

		for {
			if counter == 10 {
				break // Выход из цикла
			}

			fmt.Println("Значение счётчика:", counter)

			counter++
		}
	}

	{ // Бесконечный цикл с условием продолжения
		var counter int

		for {
			fmt.Println("Значение счётчика:", counter)

			counter++

			if counter != 10 {
				continue
			}

			break
		}
	}

	{ // for инициализация ; условие; изменение счётчика {
		for i := 0; i < 10; i++ {
			fmt.Println("Значение счётчика:", i)
		}
	}

	{ // Можно так
		i := 0

		for ; i < 10; i++ {
			fmt.Println("Значение счётчика:", i)
		}
	}

	{ // Или так
		for i := 0; i < 10; {
			fmt.Println("Значение счётчика:", i)

			i++
		}
	}

	{ // И так
		i := 0

		for i < 10 {
			fmt.Println("Значение счётчика:", i)

			i++
		}
	}

	{ // Итерация по массиву
		var numbers = [5]int{1, 2, 3, 4, 5}

		for range numbers {
			fmt.Println("Привет!")
		}
	}

	{ // Тоже самое
		for range [5]int{1, 2, 3, 4, 5} {
			fmt.Println("Привет!")
		}
	}

	{ // Тоже самое. Сахар. Выводит "Привет!" 5 раз
		for range 5 {
			fmt.Println("Привет!")
		}
	}

	{ // Можно вывести индекс переменной
		for i := range []int{1, 2, 3, 4, 5} {
			fmt.Println("Индекс:", i)
		}
	}

	{ // Можно вывести индекс и значение
		for i, v := range []int{1, 2, 3, 4, 5} {
			fmt.Println("Индекс:", i, "Значение:", v)
		}
	}

	{ // Можно вывести только значение
		for _, v := range []int{1, 2, 3, 4, 5} {
			fmt.Println("Значение:", v)
		}
	}

	{ // Сахар. С индексом начиная с 0
		for i := range 5 {
			fmt.Println("Индекс", i)
		}
	}

	{ // Так нельзя // Err: range over 5 (untyped int constant) permits only one iteration variable
		//for i, v := range 5 {
		//	fmt.Println(i, v)
		//}
	}

	{ // Выход с помощью метки из вложенного цикла
	LOOP:
		for {
			for i := 0; i < 10; i++ {
				if i == 5 {
					break LOOP
				}
			}
		}
	}

	{ // Выход с помощью метки из Селекта
	loop:
		for {
			select {
			case <-time.After(time.Nanosecond):
				break loop
			}
		}
	}

	{ //

	}
}
