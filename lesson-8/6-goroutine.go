package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	for range 10 {
		go func() {
			time.Sleep(1 * time.Second)

			runtime.Gosched()

			fmt.Println("Done!")
		}()
	}

	time.Sleep(2 * time.Second)
}
