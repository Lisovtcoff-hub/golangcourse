package main

import "fmt"

// Что выведет код?

func main() {
	s := []int{1, 2, 3, 4}

	printer := func(s []int) {
		fmt.Println(s)
	}

	printer(s)

	s[0] = 0

	printer(s)
}
