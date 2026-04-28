package main

import (
	"fmt"
	"unsafe"
)

type user struct {
	ID   int
	Name string
	Age  int
}

func (u user) GetName() string {
	return u.Name
}

func (u user) GetAge() int {
	return u.Age
}

func main() {
	u := user{
		ID:   42,
		Name: "Alice",
		Age:  25,
	}

	name := u.GetName()
	age := u.GetAge()

	fmt.Println(name, age)

	// Паддинг (выравнивание) в структурах
	// Размер структуры зависит от порядка полей и их типов
	// Поля выравниваются по 4 или 8 байт
	type user1 struct {
		ID   int32  // 4 байта, но займёт 8
		Name string // 16 байт
		Age  int32  // 4 байта, но займёт 8
	}

	fmt.Println("Размер структуры user1:", unsafe.Sizeof(user1{})) // 32 байта

	type user2 struct {
		Name string // 16 байт
		ID   int32  // 4 байта и займёт 4 байта
		Age  int32  // 4 байта и займёт 4 байта
	}

	fmt.Println("Размер структуры user2:", unsafe.Sizeof(user2{})) // 24 байта

	type user3 struct {
		ID    int  // 8 байт
		Valid bool // 1 байт, но займёт 8 байт
	}

	fmt.Println("Размер структуры user3:", unsafe.Sizeof(user3{})) // 16 байт

	type user4 struct {
		ID    int32 // 4 байта
		Valid bool  // 1 байт, но займёт 4 байта
	}

	fmt.Println("Размер структуры user4:", unsafe.Sizeof(user4{})) // 8 байт
}
