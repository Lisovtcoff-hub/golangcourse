package main

import "fmt"

func main() {
	{ // Работа с неинициализированным каналом
		var ch chan int
		ch <- 1 // fatal error: all goroutines are asleep - deadlock!
	}

	{ // Дедлок - никто не читает
		var ch = make(chan int, 1)

		ch <- 1
		ch <- 1 // fatal error: all goroutines are asleep - deadlock!
	}

	{ // Дедлок - никто не пишет
		var ch = make(chan int, 1)

		<-ch // fatal error: all goroutines are asleep - deadlock!
	}

	{ // Закрытие закрытого канала
		var ch = make(chan int, 1)
		close(ch)
		close(ch) // panic: close of closed channel
	}

	{ // Запись в закрытый канал
		var ch = make(chan int, 1)
		close(ch)
		ch <- 1 // panic: send on closed channel
	}

	{ // Чтение из закрытого канала - OK!
		var ch = make(chan int, 1)
		close(ch)

		_, ok := <-ch
		if !ok {
			fmt.Println("Channel is closed")
		}
	}

	{ // For range закрытого канала - OK!
		var ch = make(chan int, 1)

		ch <- 1

		close(ch)

		for value := range ch {
			fmt.Println("Прочитано из канала:", value)
		}

		fmt.Println("Прочитали всё из канала и вышли из цикла")
	}

	{ // Select закрытого канала - OK!
		var ch = make(chan int, 1)
		close(ch)

		select {
		case value := <-ch:
			fmt.Println("Прочитано из канала (zero value):", value)
		default:
			fmt.Println("Канал закрыт")
		}
	}
}
