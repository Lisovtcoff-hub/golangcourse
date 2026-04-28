package main

import (
	"fmt"
)

func main() {
	c := []string{"A", "B", "D", "E"}

	b := c[1:2]

	b = append(b, "TT")

	fmt.Println(c)
	fmt.Println(b)
}
