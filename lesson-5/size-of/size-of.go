package main

import (
	"fmt"
	"unsafe"
)

// unsafe.Sizeof - показывает размер на стеке
func main() {
	fmt.Println("Полностью на стеке:")
	fmt.Println("boolean:", unsafe.Sizeof(true))
	fmt.Println("int8:", unsafe.Sizeof(int8(0)))
	fmt.Println("int16:", unsafe.Sizeof(int16(0)))
	fmt.Println("int32:", unsafe.Sizeof(int32(0)))
	fmt.Println("int64:", unsafe.Sizeof(int64(0)))
	fmt.Println("array:", unsafe.Sizeof([8]byte{}))

	fmt.Println("\nТолстые указатели:")
	fmt.Println("slice:", unsafe.Sizeof([]byte{}))
	fmt.Println("string:", unsafe.Sizeof(""))
	fmt.Println("interface:", unsafe.Sizeof(any(0)))

	fmt.Println("\nОбычные указатели:")
	fmt.Println("pointer:", unsafe.Sizeof(&struct{}{}))
	fmt.Println("map:", unsafe.Sizeof(map[string]string{}))
	fmt.Println("channel:", unsafe.Sizeof(make(chan int)))
	fmt.Println("function:", unsafe.Sizeof(func() {}))

	fmt.Println("\nПустая структура:")
	fmt.Println("struct:", unsafe.Sizeof(struct{}{}))

	fmt.Println("\nСтруктуры:")
	type User struct {
		ID     int
		Verify bool
	}

	// Порядок в памяти такой же как в структуре + выравнивание
	fmt.Println("struct only int", unsafe.Sizeof(User{}))
}
