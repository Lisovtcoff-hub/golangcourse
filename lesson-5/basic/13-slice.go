package main

import (
	"fmt"
	"unsafe"
)

func main() {
	{ // Слайс - это структура
		type slice struct {
			array unsafe.Pointer // Указатель на массив [3]string{"Tom", "Alice", "Kate"}
			len   int            // Длина слайса
			cap   int            // Ёмкость слайса
		}
	}

	{ // Zero value
		var users []string
		fmt.Println(users == nil) // true
	}

	{ // Без инициализации
		var users []string

		fmt.Println(users == nil) // true

		fmt.Println("Размер слайса на стеке:", unsafe.Sizeof(users)) // 24

		_ = users
	}

	{ // С инициализацией
		var users []string = []string{"Tom", "Alice", "Kate"}

		fmt.Println("Размер:", unsafe.Sizeof(users)) // 24

		_ = users
	}

	{ // Доступ                0       1        2
		users := []string{"Tom", "Alice", "Kate"}

		// Доступ к 3 элементу
		kate := users[2]

		// Перезапись 3 элемента
		users[2] = kate

		// Перебор (есть риск выйти за границы)
		for i := 0; i < len(users); i++ {
			fmt.Println(users[i])
		}

		// Перебор (нет риска выйти за границы)
		for _, value := range users {
			fmt.Println(value)
		}
	}

	{ // Length и Capacity
		var users []string

		fmt.Println(len(users), cap(users)) // 0 0

		users = append(users, "Tom")
		fmt.Println(len(users), cap(users)) // 1 1

		users = append(users, "Tom")
		fmt.Println(len(users), cap(users)) // 2 2

		users = append(users, "Tom")
		fmt.Println(len(users), cap(users)) // 3 4

		users = append(users, "Tom")
		fmt.Println(len(users), cap(users)) // 4 4

		users = append(users, "Tom")
		fmt.Println(len(users), cap(users)) // 5 8 [8]string{"Tom", "Tom", "Tom", "Tom", "Tom", "", "", ""}
	}

	// https://go.dev/play/p/O29sYrevk28
	{ // Измеряем capacity
		var capacity int

		for range 40 {
			s := make([]int, capacity)
			s = append(s, 1)

			fmt.Println(capacity, "->", cap(s), fmt.Sprintf("%.2f", float32(cap(s))/float32(capacity)))

			capacity = cap(s)
		}
	}

	// Output
	// Рост capacity у слайса:
	// 0		->	 1
	// 1		->	 2		 2.00
	// 2		->	 4		 2.00
	// 4		->	 8		 2.00
	// 8		->	 16		 2.00
	// 16		->	 32		 2.00
	// 32		->	 64		 2.00
	// 64		->	 128	 2.00
	// 128		->	 256	 2.00
	// 256		->	 512	 2.00
	// 512		->	 848	 1.66
	// 848		->	 1280	 1.51
	// 1280		->	 1792	 1.40
	// 1792		->	 2560	 1.43
	// 2560		->	 3408	 1.33
	// 3408		->	 5120	 1.50
	// 5120		->	 7168	 1.40
	// 7168		->	 9216	 1.29
	// 9216		->	 12288	 1.33
	// 12288	->	 16384	 1.33
	// 16384	->	 21504	 1.31
	// 21504	->	 27648	 1.29
	// 27648	->	 34816	 1.26
	// 34816	->	 44032	 1.26
	// 44032	->	 55296	 1.26
	// 55296	->	 69632	 1.26
	// 69632	->	 88064	 1.26
	// 88064	->	 110592	 1.26
	// 110592	->	 139264	 1.26
	// 139264	->	 175104	 1.26
	// 175104	->	 219136	 1.25 // далее 1.25

	{ // Преаллокация через make (оптимизация для уменьшения аллокаций)
		// И длина и емкость будут 3.
		// Слайс будет проинициализирован 3 нулевыми значениями
		users := make([]string, 3)

		fmt.Println(len(users)) // 3
		fmt.Println(cap(users)) // 3

		users[0] = "Tom"
		users[1] = "Alice"
		users[2] = "Bob"

		// Длина будет 0, а емкость 3
		users2 := make([]string, 0, 3)

		fmt.Println(len(users)) // 0
		fmt.Println(cap(users)) // 3

		// Так обратиться уже нельзя, будет паника из-за выхода за границы слайса
		users[0] = "Tom"

		// Но можно добавлять элементы через append
		users2 = append(users2, "Tom")
		users2 = append(users2, "Alice")
		users2 = append(users2, "Bob")

		// Создание пустого слайса с преаллоцированым массивом на 1_000_000 элементов
		users3 := make([]string, 0, 1_000_000) // под капотом массив [1_000_000]string

		fmt.Println("users3: ", len(users3), cap(users3)) // 0 1_000_000

		// Теперь можем добавить 1_000_000 элементов без аллокаций
		for i := 0; i < 1_000_000; i++ {
			users3 = append(users3, "Tom")
		}

		fmt.Println("users3: ", len(users3), cap(users3)) // 1_000_000 1_000_000

		// При добавлении 1_000_001, произойдёт аллокация, новый массив будет больше примерно в 1.25 раз
		users3 = append(users3, "Tom")

		fmt.Println("users3: ", len(users3), cap(users3)) // 1_000_001 1_250_304

		_ = users3
	}

	{ //                    0       1       2       3      4      5	       6        7
		users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}

		users1 := users[2:6] // с 3-го по 6-й
		users2 := users[:4]  // с 1-го по 4-й
		users3 := users[3:]  // с 4-го до конца
		users4 := users[:]   // взять всё (так можно из массива сделать слайс)

		fmt.Println(users1) // ["Kate", "Sam", "Tom", "Paul"]
		fmt.Println(users2) // ["Bob", "Alice", "Kate", "Sam"]
		fmt.Println(users3) // ["Sam", "Tom", "Paul", "Mike", "Robert"]
		fmt.Println(users4) // ["Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"]
	}

	{ // Удаление           0        1       2
		users := []string{"Tom", "Alice", "Kate"}

		// Удаление по индексу
		users = append(users[:1], users[2:]...)

		fmt.Println(users) // ["Tom", "Kate"]
	}

	{
		var nums []int

		fmt.Println("nums == nil", nums == nil) // true

		nums2 := []int{}

		fmt.Println("nums2 == nil", nums2 == nil) // false

		fmt.Println(nums, nums2)
	}
}
