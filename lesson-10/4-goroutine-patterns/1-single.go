package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func worker(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Начало работы: %s\n", name)
	time.Sleep(time.Second)
	fmt.Printf("Завершение работы: %s\n", name)
}

func main() {
	var wg sync.WaitGroup

	for n := range 1 {
		wg.Add(1)
		go worker("Горутин"+strconv.Itoa(n), &wg)
	}

	wg.Wait() // ждём, пока все горутины завершатся

	fmt.Println("Все горутины завершены. Завершение программы.")
}
