package main

import (
	"fmt"
	"sync"
)

func main() {
	//runtime.GOMAXPROCS(1) // разблокировать и посмотреть поведение

	ch := make(chan int, 20)
	wg := sync.WaitGroup{}
	wg.Add(10)

	// Запуск 10 воркеров
	for i := range 10 {
		go func() {
			defer wg.Done()

			for value := range ch {
				fmt.Println("Worker", i, "Value:", value)

				//runtime.Gosched()
				//time.Sleep(1000 * time.Nanosecond)
			}
		}()
	}

	// Отправка данных в канал
	for i := range 20 {
		ch <- i
	}

	close(ch)

	wg.Wait()
}
