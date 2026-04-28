package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Руны
	// Руна — это код символа Unicode, представленный в кодировке UTF-8
	g := 'G' // Руна буквы G
	f := 'Ф' // Руна буквы Ф

	fmt.Println(utf8.RuneLen(g)) // 1
	fmt.Println(utf8.RuneLen(f)) // 2

	runes := []rune{'п', 'р', 'и', 'в', 'е', 'т', ',', ' ', 'м', 'и', 'р', '!'}

	for _, r := range runes {
		fmt.Printf("%c", r) // привет, мир!
	}

	// Руны легко проводятся к строке и обратно
	strings := string(runes)
	runes = []rune(strings)
	strings = string(runes)

	fmt.Println(strings)

	str := "Hello, world!"
	fmt.Printf("Длина строки в байтах: %d\n", len(str))                   // 13
	fmt.Printf("Длина строки в рунах: %d\n", utf8.RuneCountInString(str)) // 13

	fmt.Println("Руны в строке:")
	for i, r := range str {
		fmt.Printf("%d: %c\n", i, r)
	}

	str = "Привет, мир!"
	fmt.Printf("Длина строки в байтах: %d\n", len(str))                   // 21
	fmt.Printf("Длина строки в рунах: %d\n", utf8.RuneCountInString(str)) // 12

	fmt.Println("Руны в строке:")
	for i, r := range str {
		fmt.Printf("%d: %c\n", i, r)
	}
}
