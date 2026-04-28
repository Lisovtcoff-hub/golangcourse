package main

import "fmt"

// Условные конструкции
func main() {
	const zero = 0

	// Условие if
	if zero == 0 {
		// true
		fmt.Println("zero is zero")
	}

	// Условие if с блоком else
	if zero == 1 {
		// false
	} else {
		// true
	}

	// Много условий
	if zero == 1 {
		// false
	} else if zero == 2 {
		// false
	} else if zero == 3 {
		// false
	} else {
		// true
	}

	// Тоже самое через Switch
	switch {
	case zero == 1:
		// false
	case zero == 2:
		// false
	case zero == 3:
		// false
	default:
		// true

	}

	// Switch
	switch zero {
	case 1:
		// false
	case 2:
		// false
	case 3:
		// false
	default:
		// true
	}

	// Switch с несколькими условиями
	switch zero {
	case 0:
		// true
	case 1, 2, 3, 10_000, 1_000_000:
		// false
	default:
		// false
	}

	//	If с выражением внутри
	if ok := isZero(zero); !ok {
		fmt.Println("Alarm!")
	}
}

func isZero(i int) bool {
	return i == 0
}
