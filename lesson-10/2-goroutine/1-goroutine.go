package main

import (
	"fmt"
	"runtime"
	"time"
)

func hello() {
	fmt.Println("Hello 1")
}

func main() {
	runtime.GOMAXPROCS(4) // Количество P, которые будут использоваться для выполнения горутин
	runtime.Gosched()     // Передача управления другим горутинам

	{ // Запуск функции в горутине
		go hello()

		time.Sleep(time.Second)
	}

	{ // Запуск анонимной функции в горутине
		go func() {
			fmt.Println("Hello 2")
		}()

		time.Sleep(time.Second)
	}

	{ // То же самое
		fn := func() {
			fmt.Println("Hello 3")
		}

		go fn()

		time.Sleep(time.Second)
	}
}
