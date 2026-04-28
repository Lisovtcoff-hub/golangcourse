package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	{ // Одновременный запуск медленных операций
		var wg sync.WaitGroup

		wg.Add(3)

		now := time.Now()

		go func() {
			defer wg.Done()
			time.Sleep(time.Second) // Service A
		}()

		go func() {
			defer wg.Done()
			time.Sleep(time.Second) // Service B
		}()

		go func() {
			defer wg.Done()
			time.Sleep(time.Second) // Service C
		}()

		wg.Wait()

		fmt.Println("Время выполнения:", time.Since(now))
	}
}
