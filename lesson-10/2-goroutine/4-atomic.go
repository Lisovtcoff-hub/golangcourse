package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	{ // Atomic
		var counter atomic.Int64

		for range 10_000 {
			go func() {
				counter.Add(1)
			}()
		}

		time.Sleep(time.Second)

		fmt.Println(counter.Load())
	}
}
