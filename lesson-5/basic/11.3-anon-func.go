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
		var x int

		go func() {
			fmt.Println("Hello", x)
		}()
	}

	{ // Анонимная асинхронная функция
		var x struct{ First int }

		fn := func(x struct{ First int }) {
			fmt.Println("Hello")
		}

		go fn(x)
	}

	{ // Замыкание
		var sum int
		var i string

		count := func() {
			sum++
			i = "Hello"
		}

		count()
		fmt.Println(i) // Hello
		count()
		count()

		fmt.Println(sum) // 3
	}

}
