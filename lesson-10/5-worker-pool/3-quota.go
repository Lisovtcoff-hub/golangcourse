package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	tasks := make(chan int, 100)

	quota := make(chan struct{}, 3)

	for range 10 {
		go func() {
			for task := range tasks {
				quota <- struct{}{}

				time.Sleep(time.Second) // hard work | parallelism 3
				fmt.Println(task)

				<-quota
			}
		}()
	}

	for i := 0; i < 100; i++ {
		tasks <- i
	}

	time.Sleep(time.Minute)
}

// Example
var quota = make(chan struct{}, 50)

func HandlerHTTP(w http.ResponseWriter, _ *http.Request) {
	select {
	case quota <- struct{}{}:
	default:
		http.Error(w, "too many requests", http.StatusTooManyRequests)
	}

	time.Sleep(time.Second) // hard work | parallelism 50

	<-quota
}
