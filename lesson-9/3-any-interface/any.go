package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Пустой интерфейс в Go (interface{} или any) может содержать значение любого типа

	{ // Пустой интерфейс
		var i interface{} // то же самое что и any

		fmt.Println("Размер пустого интерфейса на стеке:", unsafe.Sizeof(i)) // 16

		i = 42
		fmt.Printf("i имеет значение %v и тип %T\n", i, i)

		i = "hello"
		fmt.Printf("i имеет значение %v и тип %T\n", i, i)

		i = true
		fmt.Printf("i имеет значение %v и тип %T\n", i, i)
	}

	{ // Функция принимает пустой интерфейс
		fn := func(v interface{}) {
			fmt.Printf("Значение: %v, Тип: %T\n", v, v)
		}

		fn(42)
		fn("hello")
		fn(true)
	}

	{ // С помощью any можно превратить Go в динамически типизированный язык
		values := []any{42, "hello", true}

		for _, v := range values {
			fmt.Printf("Значение: %v, Тип: %T\n", v, v)
		}
	}

	{ // Приведение пустого интерфейса к конкретному типу
		var i any = "hello"

		// Приведение типа с паникой при неудаче
		s := i.(string)
		fmt.Printf("Строка: %s\n", s)

		// Приведение типа
		s, ok := i.(string)
		if ok {
			fmt.Printf("Строка: %s\n", s)
		} else {
			fmt.Println("Не строка")
		}
	}

	{ // Type switch
		var i any = "hello"

		switch v := i.(type) {
		case int:
			fmt.Printf("Это число: %d\n", v)
		case string:
			fmt.Printf("Это строка: %s\n", v)
		case error:
			fmt.Printf("Это ошибка: %v\n", v)
		default:
			fmt.Println("Неизвестный тип")
		}
	}
}
