package main

import "fmt"

func main() {
	{ // Анонимная функция
		f := func(x, y int) int {
			return x + y
		}

		result := f(3, 4)

		fmt.Println(result) // 7
	}

	{ // Анонимная асинхронная функция
		go func() {
			fmt.Println("Hello")
		}()
	}

	{ // Замыкание
		var sum int

		count := func() {
			sum++
		}

		count()
		count()
		count()

		fmt.Println(sum) // 3
	}

}
