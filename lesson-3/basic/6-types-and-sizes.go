package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var ( // Целые числа
		i   int   // 8 байт
		i64 int64 // 8 байт
		i32 int32 // 4 байта
		i16 int16 // 2 байта
		i8  int8  // 1 байт

		ui   uint   // 8 байт
		ui64 uint64 // 8 байт
		ui32 uint32 // 4 байта
		ui16 uint16 // 2 байта
		ui8  uint8  // 1 байт

		b byte // 1 байт. Псевдоним для uint8. Используется в []byte
		r rune // 4 байта. Псевдоним для int32. Представляет собой код Unicode (rune - это не символ, а код символа)
	)

	fmt.Println("Размер int:", unsafe.Sizeof(i))
	fmt.Println("Размер int64:", unsafe.Sizeof(i64))
	fmt.Println("Размер int32:", unsafe.Sizeof(i32))
	fmt.Println("Размер int16:", unsafe.Sizeof(i16))
	fmt.Println("Размер int8:", unsafe.Sizeof(i8))

	fmt.Println("Размер uint:", unsafe.Sizeof(ui))
	fmt.Println("Размер uint64:", unsafe.Sizeof(ui64))
	fmt.Println("Размер uint32:", unsafe.Sizeof(ui32))
	fmt.Println("Размер uint16:", unsafe.Sizeof(ui16))
	fmt.Println("Размер uint8:", unsafe.Sizeof(ui8))

	fmt.Println("Размер байта:", unsafe.Sizeof(b))
	fmt.Println("Размер руны:", unsafe.Sizeof(r))

	var ( // Числа с плавающей точкой
		f64 float64 // 8 байт
		f32 float32 // 4 байта
	)

	fmt.Println("Размер float64:", unsafe.Sizeof(f64))
	fmt.Println("Размер float32:", unsafe.Sizeof(f32))

	var ( // Булевы значения
		b1 bool // 1 байт. Может быть: true или false
	)

	fmt.Println("Размер булевого значения:", unsafe.Sizeof(b1))

	var ( // Строка
		s string // Размер 16 байт
	)

	fmt.Println("Размер строки:", unsafe.Sizeof(s))

	// Почему строка 16 байт? Строка - это структура:
	type string struct {
		Data uintptr // 8 байт. Указатель на массив байт: []byte
		Len  int     // 8 байт. Длина строки
	}
}
