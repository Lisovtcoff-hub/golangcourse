package main

import (
	"fmt"
	"sync"
	"unsafe"
)

func main() {
	{ // Zero value
		var m map[string]int

		fmt.Println(m == nil) // true
	}

	{ // Без инициализации   map[key]value
		var people map[string]int // указатель unsafe.Pointer на структуру мапы в куче

		fmt.Println("Размер мапы на стеке:", unsafe.Sizeof(people)) // 8

		fmt.Println(people == nil) // true

		_ = people
	}

	{ // Инициализация
		var people = map[string]int{
			"Bob":   2,
			"Tom":   1,
			"Sam":   4,
			"Alice": 8,
		}

		fmt.Println(people) // map[Tom:1 Bob:2 Sam:4 Alice:8]
	}

	{ // Доступ
		var people = map[string]int{
			"Tom":   1,
			"Bob":   2,
			"Sam":   4,
			"Alice": 0,
		}

		fmt.Println(people["Alice"]) // 8
		fmt.Println(people["Bob"])   // 2
		people["Bob"] = 32
		fmt.Println(people["Bob"]) // 32
	}

	{ // Проверка значения
		var people = map[string]int{
			"Tom":   0,
			"Bob":   2,
			"Sam":   4,
			"Alice": 8,
		}

		val, ok := people["Tom"]
		if ok {
			fmt.Println("VALUE!", val)
		}
	}

	{ // Добавление
		var people = map[string]int{"Tom": 1, "Bob": 2}
		people["Kate"] = 128
		fmt.Println(people) // map[Tom:1  Bob:2  Kate:128]
	}

	{ // Удаление
		var people = map[string]int{"Tom": 1, "Bob": 2, "Sam": 8}
		delete(people, "Bob")
		fmt.Println(people) // map[Tom:1  Sam:8]
	}

	{ // Итерация
		var people = map[string]int{
			"Tom": 1,
			"Bob": 2,
			"Sam": 4,
		}

		fmt.Println("Iterate!")
		for key, value := range people {
			fmt.Println(key, value) // В случайном порядке
		}
	}

	{ // Ключами могут быть типы, которые поддерживают операторы сравнения == и !=
		// Булев               true == false, true != false
		// Целочисленный       5 == 10, 5 != 10
		// С плавающей точкой  3.14 == 2.71, 3.14 != 2.71
		// Строка              "hello" == "world", "a" != "b"
		// Указатель           p1 == p2, p1 != p2
		// Канал               ch1 == ch2, ch1 != ch2
		// Интерфейс           i1 == i2, i1 != i2
		// Массив              [3]int{1, 2, 3} == [3]int{1, 2, 3}
	}

	{ // Сет
		var people = map[string]struct{}{}

		people["Tom"] = struct{}{}
		people["Bob"] = struct{}{}
		people["Sam"] = struct{}{}
		people["Sam"] = struct{}{}
		people["Sam"] = struct{}{}

		fmt.Println(people) // map[Bob:{} Tom:{} Sam:{}]
	}

	{ // Защита мьютексом
		type Cache struct {
			mx sync.Mutex
			m  map[string]string
		}

		cache := Cache{
			m: map[string]string{},
		}

		cache.mx.Lock()
		cache.m["key"] = "value"
		cache.mx.Unlock()
	}

	{ // RW мьютекс
		type Cache struct {
			mx sync.RWMutex
			m  map[string]string

			s string
		}

		cache := Cache{
			m: map[string]string{},
		}

		// Write
		cache.mx.Lock()
		cache.m["key"] = "value"
		cache.mx.Unlock()

		// Read
		cache.mx.RLock()
		_ = cache.m["key"]
		cache.mx.RUnlock()
	}
}
