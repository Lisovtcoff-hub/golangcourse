package main

import "fmt"

func main() {
	var i int

	closure := func() {
		i++
	}

	closure()
	closure()
	fmt.Println(i)
}
