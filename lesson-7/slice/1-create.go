package main

import (
	"fmt"
	"unsafe"
)

// На самом деле []int это структура:
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {
	// Неинициализированный слайс
	var nilSlice []string

	fmt.Println("nilSlice == nil:", nilSlice == nil) // true
	fmt.Println("len", len(nilSlice))                // 0
	fmt.Println("cap", cap(nilSlice))                // 0

	// Инициализированный слайс
	var initSlice = []string{}

	fmt.Println("initSlice == nil:", initSlice == nil) // false
	fmt.Println("len", len(initSlice))                 // 0
	fmt.Println("cap", cap(initSlice))                 // 0

	// Если len == 0, нельзя обратиться к 1 элементу по индексу
	//fmt.Println("Read", initSlice[0]) // panic: index out of range

	// Записать тоже нельзя
	//initSlice[0] = "Alice" // panic: index out of range

	// Можно добавить элемент в конец слайса
	initSlice = append(initSlice, "Alice")

	fmt.Println("initSlice after append:", initSlice) // [Alice]
	fmt.Println("len", len(initSlice))                // 1
	fmt.Println("cap", cap(initSlice))                // 1

	// Теперь можем обратиться к 1 элементу по индексу
	fmt.Println("Read", initSlice[0])

	// И можем перезаписать
	initSlice[0] = "Bob"

	// Инициализированный через литерал слайс сразу с элементами
	var users = []string{"Bob", "Tom", "Alice"}
	_ = users

	// Создание слайса с помощью make
	var makeSlice = make([]string, 0, 5)
	fmt.Println("makeSlice == nil:", makeSlice == nil) // false
	fmt.Println("len", len(makeSlice))                 // 0
	fmt.Println("cap", cap(makeSlice))                 // 5
}
