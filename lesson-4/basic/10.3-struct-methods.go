package main

import "fmt"

// Пример встраивания структур с методами

type Apple struct{}

func (a Apple) Print() {
	fmt.Println("I'm an apple")
}

type Banana struct {
	Apple
}

//func (b Banana) Print() {
//	fmt.Println("I'm a banana")
//}

// При встраивании методы переопределяются при совпандении имен
func main() {
	apple := Apple{}

	banana := Banana{
		Apple: apple,
	}

	banana.Print() // I'm an apple
}
