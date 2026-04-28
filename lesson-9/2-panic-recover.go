package main

import "fmt"

func mayPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Восстановились после паники:", r)
		}
	}()

	fmt.Println("Перед паникой")

	panic("Что то пошло не так!")

	fmt.Println("Это никогда не будет напечатано")
}

func main() {
	mayPanic()

	fmt.Println("Программу удалось спасти от паники")
}
