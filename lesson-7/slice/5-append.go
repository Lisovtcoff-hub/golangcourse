package main

import "fmt"

func main() {
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
}
