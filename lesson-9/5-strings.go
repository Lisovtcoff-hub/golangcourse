package main

import (
	"fmt"
	"strings"
)

func main() {
	// Сравнение строк
	str1 := "Hello, World!"
	str2 := "hello, world!"
	fmt.Println(strings.EqualFold(str1, str2)) // true (сравнение без учета регистра)

	// Преобразование регистра
	str := "Hello, World!"
	fmt.Println(strings.ToLower(str)) // "hello, world!"
	fmt.Println(strings.ToUpper(str)) // "HELLO, WORLD!"

	// Проверка наличия подстроки
	substr := "World"
	fmt.Println(strings.Contains(str, substr)) // true

	// Разделение строки
	str = "a,b,c,d,e"
	parts := strings.Split(str, ",")
	fmt.Println(parts) // [a b c d e]

	// Объединение строк
	parts = []string{"a", "b", "c", "d", "e"}
	str = strings.Join(parts, ",")
	fmt.Println(str) // "a,b,c,d,e"

	// Замена подстроки
	str = "Hello, World!"
	newStr := strings.ReplaceAll(str, "World", "Go")
	fmt.Println(newStr) // "Hello, Go!"

	// Проверка префикса и суффикса
	fmt.Println(strings.HasPrefix(str, "Hello"))  // true
	fmt.Println(strings.HasSuffix(str, "World!")) // true

	// Индекс подстроки
	fmt.Println(strings.Index(str, "World")) // 7

	// Повторение строки
	fmt.Println(strings.Repeat("Go", 3)) // "GoGoGo"

	// Обрезка символов
	str = "!!!Hello, World!!!"
	fmt.Println(strings.Trim(str, "!")) // "Hello, World"

	// Обрезание символов
	fmt.Println(strings.TrimLeft(str, "!"))  // "Hello, World!!!"
	fmt.Println(strings.TrimRight(str, "!")) // "!!!Hello, World"

	// Обрезание пробелов
	str = "   Hello, World!   "
	fmt.Println(strings.TrimSpace(str)) // "Hello, World!"

	// Обрезание по функции
	str = "123Hello, World!123"
	fmt.Println(strings.TrimFunc(str, func(r rune) bool {
		return r >= '0' && r <= '9'
	})) // "Hello, World!"

	// Полное обрезание по префиксу и суффиксу
	str = "Hello, World!"
	fmt.Println(strings.TrimPrefix(str, "Hello, ")) // "World!"
	fmt.Println(strings.TrimSuffix(str, "!"))       // "Hello, World"
}
