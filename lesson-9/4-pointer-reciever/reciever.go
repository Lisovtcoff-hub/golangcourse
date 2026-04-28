package main

import "fmt"

type Interface interface {
	method()
}

func fn(v Interface) {
	v.method()
}

func main() {
	val := &T{}
	fn(val)
}

type T struct{}

func (t *T) method() {
	fmt.Printf("Hello!")
}
