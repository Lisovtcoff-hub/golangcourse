package main

import "fmt"

type MyInterface interface{}

type Example struct {
	Value string
}

func example1() MyInterface {
	var e *Example

	return e
}

func example2() MyInterface {
	return nil
}

func main() {
	e1 := example1()
	e2 := example2()

	fmt.Println("e1 == e2:", e1 == e2)

	fmt.Println("e1 == nil:", e1 == nil)
	fmt.Println("e2 == nil:", e2 == nil)

	fmt.Printf("Value e1: %v\n", e1)
	fmt.Printf("Value e2: %v\n", e2)

	fmt.Printf("Type e1: %T\n", e1)
	fmt.Printf("Type e2: %T\n", e2)
}
