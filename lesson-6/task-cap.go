package main

import "fmt"

func main() {
	users := []string{"Bob", "Alice", "Kate"}

	users2 := users[0:2] // А если users[0:2:2]?
	fmt.Println("Users2", users2)
	fmt.Println("Len", len(users2)) // 2
	fmt.Println("Cap", cap(users2)) // 3

	users2 = append(users2, "X") // Изменится ли в 1 слайсе?

	fmt.Println("Users1", users)
	fmt.Println("Users2", users2)
}
