package main

import "fmt"

// Глобальная типизированая переменная
var globalTypedVar int32 = 42

// Автовыведение типа
var globalUntypedVar = 42

// С нулевым значением
var zeroCounter int

// Локальные переменные
func main() {
	// Локальные переменные пишутся с маленькой буквы в camalCase

	// Типизированная
	var typedVar int32 = 42
	typedVar2 := int32(42)

	// Не типизированная, с автовыведением типа
	untypedVar := 42
	untypedVar2 := 42

	// С нулевым значением int
	var zeroValueInt int
	zeroValueInt2 := 0

	// В блоке var можно создавать несколько переменных
	var (
		x = 10
		y = 20
	)

	// Меняем местами значения переменных
	y, x = x, y

	// Область видимости переменной
	// ограничена блоком из фигурных скобок { ... } (Например: if, for, switch, func)
	// доступ к переменной нельзя получить за пределами блока { ... }

	// Создаем переменную с нулевым значением
	var counter int

	// Но можем переопределить
	counter = 52

	// Открытая фигурная скобка { означает новую область видимости
	{
		// Можем переопределить
		counter = 62

		// Совершенно другая переменная !!!
		counter := ""
		_ = counter

		newCounter := 100 // не видна за пределами блока
		_ = newCounter
	}

	// _ = newCounter // ошибка компиляции

	_ = counter

	shadowing()
	notShadowing()
	_, _, _, _, _, _, _, _, _ = globalTypedVar, globalUntypedVar, zeroCounter, typedVar, typedVar2, untypedVar, untypedVar2, zeroValueInt, zeroValueInt2
}

// Глобальная переменная
var shadow = 10

func shadowing() {
	// Локальная переменная, затеняющая глобальную переменную
	shadow := 20
	fmt.Println(shadow) // Выведет 20, так как локальная переменная затеняет глобальную

	if true {
		// Переменная в блоке, затеняющая переменную в функции main
		shadow := 30
		fmt.Println(shadow) // Выведет 30, так как переменная в блоке затеняет переменную в функции main
	}

	fmt.Println(shadow) // Выведет 20, так как мы вернулись к области видимости функции main
}

func notShadowing() {
	fmt.Println(shadow) // Выведет 10, так как переменная x в этой функции не затеняется
}
