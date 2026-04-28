package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Указатель - это переменная, которая хранит адрес памяти на другую переменную (т.е. указывает на другую переменную)
	// Указатель может быть nil (не указывать ни на какую переменную)

	// * - разыменовать что угодно
	// & - взять указатель на что угодно
	{
		var p *int

		fmt.Println(p == nil) // true

		//fmt.Println(*p) // panic: nil pointer dereference

		x := 42

		p = &x // Взятие адреса переменной

		_ = *p // Разыменование указателя (получение значения по адресу)

		fmt.Println("p:", p)  // 0xc000012108
		fmt.Println("*p", *p) // 42
	}

	{
		// *int - тип указателя
		// &x - взятие указателя

		var x = 4
		var p = &x
		_ = p
	}

	{ // Указатель может указывать на указатель
		x := 42

		i := &x // Взятие адреса переменной

		j := &i

		k := &j

		fmt.Println("Размер указателя i:", unsafe.Sizeof(i)) // 8

		fmt.Println("k:", k)       // 0xc00008c050
		fmt.Println("*k:", *k)     // 0xc00008c048
		fmt.Println("**k:", **k)   // 0xc000012120
		fmt.Println("***k:", ***k) // 42

		fmt.Println("*k == j", *k == j)       // true
		fmt.Println("**k == i", **k == i)     // true
		fmt.Println("***k == 42", ***k == 42) // true
	}

	{ // nil pointers разных типов не сравниваются
		var x *int
		var y *string

		// fmt.Println(x == y) // true // Compile error
		_, _ = x, y
	}

	{ // Указатель на string
		s := "Hello"
		p := &s
		_ = p
	}

	{ // Указатель на struct
		type User struct {
			Name string
		}

		s := User{Name: "Tom"}
		p := &s

		//	Для указателя на структуру можно использовать сокращенную запись
		p = &User{Name: "Tom"}
		_ = p

		// Или даже так (устаревшее)
		p = new(User)
		_ = p
	}

	{ // Передача в функцию
		x := 13

		pointer := wantsPointer(&x)
		if pointer == nil {
			panic("pointer is not nil")
		}

		fmt.Println(*pointer)
	}

}

func wantsPointer(pointer *int) *int {
	fmt.Println("*pointer:", *pointer)

	return pointer
}
