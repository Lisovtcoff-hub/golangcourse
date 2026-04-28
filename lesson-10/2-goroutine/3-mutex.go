package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	{ // Mutex
		mx := sync.Mutex{}
		var counter int

		for range 10_000 {
			go func() {
				mx.Lock() // <- Блокировка
				counter++
				mx.Unlock()
			}()
		}

		time.Sleep(time.Second)

		fmt.Println(counter)
	}

}
