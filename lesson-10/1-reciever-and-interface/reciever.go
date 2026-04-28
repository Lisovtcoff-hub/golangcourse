package main

import "fmt"

type Interface interface {
	method()
}

func fn(v Interface) {
	v.method()
}

func main() {
	t := &T{}
	fn(t)

	fmt.Println("Payload:", t.Payload)
}

type T struct {
	Payload string
}

func (t *T) method() {
	t.Payload = "Value!"
	fmt.Println("Method called")
}
