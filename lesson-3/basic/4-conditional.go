package main

// Операторы сравнения
func main() {
	// Результат сравнения - булево значение
	// true или false

	black, white := "black", "white"

	// == равно
	if black == white {
		// false
	}

	// != не равно
	if black != white {
		// true
	}

	zero, one := 0, 1

	// < меньше
	if zero < one {
		// true
	}

	// > больше
	if zero > one {
		// false
	}

	// <= меньше или равно
	if zero <= one {
		// true
	}

	// >= больше или равно
	if zero >= one {
		// false
	}

	// Логические операторы

	// && логическое И
	_ = true && true   // true
	_ = true && false  // false
	_ = false && true  // false
	_ = false && false // false

	// || логическое ИЛИ
	_ = true || true   // true
	_ = true || false  // true
	_ = false || true  // true
	_ = false || false // false

	// ! логическое НЕ (использовать редко и аккуратно, ухудшает читаемость)
	_ = !true  // false
	_ = !false // true
}
