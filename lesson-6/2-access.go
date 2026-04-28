package main

import (
	"fmt"
)

func main() {
	{ // Доступ             0       1        2
		users := []string{"Tom", "Alice", "Kate"}

		fmt.Println(len(users)) // 3
		fmt.Println(cap(users)) // 3

		// Доступ к 3 элементу
		kate := users[2]

		// Перезапись 3 элемента
		users[2] = kate

		// Перебор (есть риск выйти за границы)
		for i := 0; i < len(users); i++ {
			fmt.Println(users[i])
		}

		// Перебор (нет риска выйти за границы)
		for _, user := range users {
			fmt.Println(user)
		}

		//	Добавить элемент в конец слайса
		users = append(users, "Paul")

		// Clear - сброс значений слайса в zero value, длина и ёмкость остаются прежними
		clear(users)

		fmt.Println("After clear:", users) // ["", "", "", ""]

		// Создание слайса с указанием длины и ёмкости
		users2 := make([]string, 3)    // length 3, capacity 3
		users3 := make([]string, 0, 3) // length 0, capacity 3

		_, _ = users2, users3
	}

	{ // Копирование
		src := []string{"Tom", "Alice", "Kate"}
		dst := make([]string, len(src))

		copy(dst, src) // Копирование слайса

		fmt.Println("Copy", dst) // ["Tom", "Alice", "Kate"]
	}

	{ // Из строки можно скопировать в []byte, но не наоборот, т.к. строка неизменяемая
		users := []byte("Tom")

		copy(users, "Alice") // Копирование слайса

		fmt.Println("Copy string", string(users)) // "Ali"
	}
}
