package main

import (
	"fmt"
	"unsafe"
)

type AddFunc func(int, int) int

type IncrementFunc func(int) int

func add(x int, y int) int {
	return x + y
}

func inc(x int) int {
	return x + 1
}

func main() {
	{ // Zero value функции
		var f AddFunc
		fmt.Println(f == nil) // true
	}

	{ // Функция как тип
		var f = add

		fmt.Println("Размер фунции (как объекта первого класса) на стеке:", unsafe.Sizeof(f)) // 8 (pointer)

		result := f(3, 4)

		fmt.Println(result) // 7
	}

	{ // Функция как тип
		var i IncrementFunc = inc

		result := i(1)

		fmt.Println(result) // 2
	}

	{ // Функция как объект первого класса
		f := Foo(add)

		result := f(3)

		fmt.Println(result) // 4
	}

}

// Foo принимает функцию AddFunc и возвращает функцию IncrementFunc
func Foo(f AddFunc) IncrementFunc {
	return func(x int) int {
		return f(x, 1)
	}
}
