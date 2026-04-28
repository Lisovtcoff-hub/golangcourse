package main

import (
	"fmt"
	"time"
)

// У команды go есть флаг -race, который позволяет найти гонки данных. Обычно запускают в тестах. Но можно и в go run:
// CGO_ENABLED=1 go run -race ./2-data-race-counter.go
func main() {
	//runtime.GOMAXPROCS(1) // Количество P

	{ // Гонка данных
		var counter int

		for range 10_000 {
			go func() {
				counter++
			}()
		}

		time.Sleep(time.Second)

		fmt.Println(counter)
	}
}
