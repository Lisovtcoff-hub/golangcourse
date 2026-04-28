package main

import (
	"fmt"
	"sync"
)

func Init() {
	fmt.Println("Инициализация выполнена")
}

func main() {
	// используется для того, чтобы гарантировать, что определённый код выполнится только один раз,
	// даже если несколько горутин попытаются выполнить его одновременно

	var wg sync.WaitGroup
	var once sync.Once

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			once.Do(Init)

			fmt.Println("Горутина завершена")
		}()
	}

	wg.Wait()
}
