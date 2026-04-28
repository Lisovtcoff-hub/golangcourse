package main

import "fmt"

func main() {
	{ // Slicing            0       1       2       3      4      5	       6        7
		users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}

		fmt.Println(users[:])   // ["Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"]
		fmt.Println(users[0:0]) // []
		fmt.Println(users[1:1]) // []
		fmt.Println(users[:1])  // ["Bob"]
		fmt.Println(users[0:1]) // ["Bob"]
		fmt.Println(users[1:2]) // ["Alice"]
		fmt.Println(users[2:4]) // ["Kate", "Sam"]
		fmt.Println(users[3:6]) // ["Sam", "Tom", "Paul
		fmt.Println(users[7:])  // ["Robert"]
		fmt.Println(users[7:8]) // ["Robert"]
		fmt.Println(users[8:8]) // []
		fmt.Println(users[9:])  // panic: slice bounds out of range
		fmt.Println(users[:9])  // panic: slice bounds out of range
	}

	{ // Slicing            0       1       2       3      4      5	       6        7
		users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}

		users1 := users[2:6]
		users2 := users[:4]
		users3 := users[3:]
		users4 := users[:] // взять всё (так можно из массива сделать слайс)

		fmt.Println(users1) // ["Kate", "Sam", "Tom", "Paul"]
		fmt.Println(users2) // ["Bob", "Alice", "Kate", "Sam"]
		fmt.Println(users3) // ["Sam", "Tom", "Paul", "Mike", "Robert"]
		fmt.Println(users4) // ["Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"]
	}

	{ // А ещё можно ограничить ёмкость слайса
		users := []string{"Bob", "Alice", "Kate"}
		users1 := users[0:2]            // берем первые 2 элемента
		fmt.Println("Len", len(users1)) // 2
		fmt.Println("Cap", cap(users1)) // 3 <- ёмкость массива под капотом

		users2 := users[0:2:2]          // <- ограничиваем ёмкость
		fmt.Println("Len", len(users2)) // 2
		fmt.Println("Cap", cap(users2)) // 2 <- ограничили ёмкость слайса (массив под капотом не изменился)
	}

	{ // Удаление           0        1       2
		users := []string{"Tom", "Alice", "Kate"}

		// Удаление по индексу
		users = append(users[:1], users[2:]...)

		fmt.Println(users) // ["Tom", "Kate"]
	}
}
