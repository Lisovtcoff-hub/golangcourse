package main

import "fmt"

// Что выведет код?

func main() {
	a := 1
	p := &a
	b := &p

	*p = 3
	**b = 5

	a += 4 + *p + **b

	fmt.Println("Result:", *p)
}
