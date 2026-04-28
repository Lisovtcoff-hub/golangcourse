package main

import (
	"fmt"
	"math"
)

type Alice struct {
	Name string
}

var alice = Alice{Name: "Alice"}

func main() {
	// Подробнее смотрите https://pkg.go.dev/fmt

	// Популярное
	fmt.Printf("%v\n", alice)   // универсальный формат
	fmt.Printf("%d\n", 42)      // 42 - целое число
	fmt.Printf("%s\n", "hello") // hello - строка
	fmt.Printf("%q\n", "hello") // "hello" - строка в кавычках
	fmt.Printf("%t\n", true)    // true - булево значение
	fmt.Printf("%%\n")          // % - экранирование символа %

	// Структура
	fmt.Printf("%T\n", alice)  // main.Alice - пакет.Тип
	fmt.Printf("%v\n", alice)  // {Alice} - значение
	fmt.Printf("%+v\n", alice) // {Name:Alice} - поле:значение
	fmt.Printf("%#v\n", alice) // main.Alice{Name:"Alice"} - пакет.Тип{поля}

	// Слайс
	fmt.Printf("%T\n", []int{1, 2, 3})  // []int - тип
	fmt.Printf("%v\n", []int{1, 2, 3})  // [1 2 3] - значение
	fmt.Printf("%#v\n", []int{1, 2, 3}) // []int{1, 2, 3} - значение с типом
	fmt.Printf("%p\n", []int{1, 2, 3})  // 0xc0000b6010 - адрес в памяти

	// Мапа
	fmt.Printf("%T\n", map[string]int{"one": 1})  // map[string]int - тип
	fmt.Printf("%v\n", map[string]int{"one": 1})  // map[one:1] - значение
	fmt.Printf("%#v\n", map[string]int{"one": 1}) // map[string]int{"one":1} - значение с типом
	fmt.Printf("%p\n", map[string]int{"one": 1})  // 0xc0000b6010 - адрес в памяти

	// Целое число
	fmt.Printf("%d\n", 42)   // 42
	fmt.Printf("%T\n", 42)   // int - тип
	fmt.Printf("%+d\n", 42)  // +42 - знак
	fmt.Printf("%+d\n", -42) // -42 - знак
	fmt.Printf("%-5d\n", 42) // "42   " - выравнивание по левому краю
	fmt.Printf("%5d\n", 42)  // "   42" - выравнивание по правому краю
	fmt.Printf("%05d\n", 42) // "00042" - заполнение нулями
	fmt.Printf("%b\n", 42)   // 101010 - двоичное представление
	fmt.Printf("%#b\n", 42)  // 0b101010 - двоичное представление с префиксом 0b
	fmt.Printf("%o\n", 42)   // 52 - восьмеричное представление
	fmt.Printf("%O\n", 42)   // 0o52 - восьмеричное представление в верхнем регистре
	fmt.Printf("%#x\n", 42)  // 0x2a - шестнадцатеричное представление
	fmt.Printf("%#X\n", 42)  // 0X2A - шестнадцатеричное представление в верхнем регистре

	// Строка
	fmt.Printf("%s\n", "hello")  // hello
	fmt.Printf("%T\n", "hello")  // string - тип
	fmt.Printf("%q\n", "hello")  // "hello" - кавычки
	fmt.Printf("%#q\n", "hello") // `hello` - обратные кавычки
	fmt.Printf("%x\n", "hello")  // 68656c6c6f - шестнадцатеричное представление
	fmt.Printf("%X\n", "hello")  // 68656C6C6F - шестнадцатеричное представление в верхнем регистре

	// Число с плавающей точкой
	fmt.Printf("%f\n", math.Pi)    // 3.141593
	fmt.Printf("%e\n", math.Pi)    // 3.141593e+00 - экспоненциальное представление
	fmt.Printf("%E\n", math.Pi)    // 3.141593E+00 - экспоненциальное представление в верхнем регистре
	fmt.Printf("%g\n", math.Pi)    // 3.141592653589793 - компактное представление
	fmt.Printf("%15f\n", math.Pi)  // "       3.141593" - выравнивание по правому краю
	fmt.Printf("%-15f\n", math.Pi) // "3.141593       " - выравнивание по левому краю

	// Округление чисел с плавающей точкой
	fmt.Printf("%.2f\n", math.Pi)  // 3.14 - округление: 2 знака после запятой
	fmt.Printf("%.4f\n", math.Pi)  // 3.1416 - округление: 4 знака после запятой
	fmt.Printf("%.f\n", math.Pi)   // 3 - округление: 0 знаков после запятой
	fmt.Printf("%9.2f\n", math.Pi) // "     3.14" - выравнивание и округление
	fmt.Printf("%9.f\n", math.Pi)  // "        3" - выравнивание и округление по правому краю
	fmt.Printf("%-9.f\n", math.Pi) // "3        " - выравнивание и округление по левому краю

	// Пробел
	fmt.Printf("% d\n", 42)      // " 42"
	fmt.Printf("% d\n", -42)     // "-42"
	fmt.Printf("% x\n", "hello") // "68 65 6c 6c 6f"
}
