package main

import "fmt"

// Что выведет код?

func changePointer(p *int) {
	v := 3
	p = &v
}

func main() {
	v := 5
	p := &v

	fmt.Println(*p)

	changePointer(p)

	fmt.Println(*p)
}
