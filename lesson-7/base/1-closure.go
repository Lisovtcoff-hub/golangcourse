package main

import "fmt"

type Closure struct {
	i *int
}

func (c Closure) Func() {
	*(c.i)++
}

func main() {
	var i int

	closure := Closure{
		i: &i,
	}

	closure.Func()
	closure.Func()

	fmt.Println(i)
}
