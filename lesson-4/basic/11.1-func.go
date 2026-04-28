package main

import "fmt"

// Не видна за пределами пакета
func private() {
	fmt.Println("первая буква определяет доступность. маленькая первая буква - private")
}

// Public видна за пределами пакета
func Public() {
	fmt.Println("Большая первая буква - public")
}

func main() {
	private()
	Public()

	err := MaybeError("Alice", 25, false)
	if err != nil {
		fmt.Println(err)
	}

	user, err := GetUser(42)
	if err != nil {
		fmt.Println(err)
	}

	user, err = getUser(42)
	if err != nil {
		fmt.Println(err)
	}

	user, err = neverDoThat(42)
	if err != nil {
		fmt.Println(err)
	}

	_ = user
}

var (
	ErrEmptyName   = fmt.Errorf("name is empty")
	ErrNegativeAge = fmt.Errorf("age is negative")
	ErrBlockedUser = fmt.Errorf("user is blocked")
)

// MaybeError возвращает nil или ошибку
func MaybeError(name string, age int, isBlocked bool) error {
	if name == "" {
		return ErrEmptyName
	}

	if age < 0 {
		return ErrNegativeAge
	}

	if isBlocked {
		return ErrBlockedUser
	}

	return nil
}

var ErrNotFound = fmt.Errorf("user not found")

type User struct {
	ID   int
	Name string
	Age  int
}

// GetUser возвращает структуру или ошибку
func GetUser(id int) (User, error) {
	var user User

	if id == 0 {
		return user, ErrBlockedUser
	}

	if id == 42 {
		return User{
			ID:   42,
			Name: "Alice",
			Age:  25,
		}, nil
	}

	return user, ErrNotFound
}

// Инициализация переменных в блоке с возвращаемыми значениеми
func getUser(id int) (user User, err error) {
	if id == 0 {
		return user, ErrBlockedUser
	}

	if id == 42 {
		return User{
			ID:   42,
			Name: "Alice",
			Age:  25,
		}, nil
	}

	return user, ErrNotFound
}

// Голый return никогда не используется. Если вы так сделате, вас сожгут)
func neverDoThat(id int) (user User, err error) {
	if id == 0 {
		err = ErrBlockedUser

		return
	}

	if id == 42 {
		return User{
			ID:   42,
			Name: "Alice",
			Age:  25,
		}, nil
	}

	err = ErrNotFound

	return
}
