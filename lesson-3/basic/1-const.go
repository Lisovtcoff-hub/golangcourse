package main

import "math"

// Константы
const (
	stringConst = "hello"
	intConst    = 42
	boolConst   = true
	floatConst  = 3.14
)

// Типизированные и нетипизированные константы
const (
	typedIntConst      int     = 42
	untypedIntConst            = 42
	typedFloatConst    float64 = 3.14
	untypedFloatConst          = 3.14
	typedStringConst   string  = "hello"
	untypedStringConst         = "hello"
)

// А ля e-num
const (
	_ = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// Смешанные типы данных
const (
	A = iota     // 0
	B = 3.14     // 3.14
	C = iota * 2 // 2 * 2 = 4
	D = "string" // "string"
	E = iota     // 4
)

// Смещение значений
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

// Использование в битовых масках
const (
	FlagNone = 1 << iota
	FlagRead
	FlagWrite
	FlagExecute
)

func main() {
	// Локальные константы
	const (
		StringConst = "hello"
		IntConst    = 42
		BoolConst   = true
	)

	_ = math.Pi
	_ = math.MaxInt

	_, _, _, _, _, _, _, _, _, _ = stringConst, intConst, boolConst, floatConst, typedIntConst, untypedIntConst, typedFloatConst, untypedFloatConst, typedStringConst, untypedStringConst
	_, _, _, _, _, _, _, _, _, _ = Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday, KB, MB, GB
	_, _, _, _ = FlagNone, FlagRead, FlagWrite, FlagExecute
	_, _, _, _, _ = A, B, C, D, E
	_, _, _ = StringConst, IntConst, BoolConst
}
