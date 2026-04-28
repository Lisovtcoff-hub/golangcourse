package main

import (
	"fmt"
)

// Что выведет код? Как исправить?

func process(temp *int32) {
	var value2 int32 = 200

	*temp = value2
}

func main() {
	var value1 int32 = 100
	pointer := &value1

	fmt.Println("Value:", *pointer)
	fmt.Println("Pointer:", pointer)

	process(pointer)

	fmt.Println("Value:", *pointer)
	fmt.Println("Pointer:", pointer)
}
