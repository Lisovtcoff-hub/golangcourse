package main

import "fmt"

// Мы хотим иногда делать обёртки над чужими типами
// Есть 2 варианта это сделать:

// 1. Композиция Composition - структура просто содержит в себе другую структуру
// type Banana struct {
//   Apple Apple
// }

// 2. Встраивание Embedding - при совпадении имён, поля и методы встроеного типа переопределяются. Похоже на наследование
// type Banana struct {
//   Apple
// }

func main() {
	type Apple struct {
		Payload string
	}

	type Banana struct {
		Apple Apple // Composition
		//Apple   // Embedding
		//Payload string
	}

	a := Apple{
		Payload: "I'm an apple",
	}

	b := Banana{
		Apple: a,
		//Payload: "I'm a banana",
	}

	fmt.Println(b.Apple.Payload) // I'm an apple
	//fmt.Println(b.Payload) // I'm a banana
}
