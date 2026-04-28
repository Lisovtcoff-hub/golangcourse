package main

import (
	"fmt"
	"sync"
)

func main() {
	{ // WaitGroup - это счётчик, который позволяет дождаться завершения всех горутин
		var (
			mx      sync.Mutex
			counter int
			wg      sync.WaitGroup
		)

		for range 100 {
			wg.Add(1)

			go func() {
				defer wg.Done()

				mx.Lock()
				defer mx.Unlock()

				counter++
			}()
		}

		wg.Wait()

		fmt.Println(counter)
	}
}
