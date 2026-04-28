package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Массивы хранятся в памяти последовательно

	// Массивы в Go - это набор из элементов одного типа с фиксированным размером
	var days [7]string
	days[0] = "Пн"
	days[1] = "Вт"
	days[2] = "Ср"
	days[3] = "Чт"
	days[4] = "Пт"
	days[5] = "Сб"
	days[6] = "Вс"

	// Инициализация массива с помощью литерала
	days = [7]string{"Пн", "Вт", "Ср", "Чт", "Пт", "Сб", "Вс"}

	// Размер массива - часть его типа: [7]string и [3]string - это разные типы

	// Автовыведение типа массива
	two := [...]int{1, 2}      // [2]int
	three := [...]int{1, 2, 3} // [3]int

	fmt.Println("Размер массива из 2 int:", unsafe.Sizeof(two))   // 16 байт (int = 8 байт * 2 элемента)
	fmt.Println("Размер массива из 3 int:", unsafe.Sizeof(three)) // 24 байта (int = 8 байт * 3 элемента)

	// Длина массива
	length := len(days) // 7

	// Емкость (емкость массива соответствует его длине)
	capacity := cap(days) // 7

	// Получение значения
	monday := days[0] // "Пн"

	// Изменение значения
	days[0] = "Monday"
	monday = days[0] // "Monday"

	// Последний элемент массива
	last := len(days) - 1
	sunday := days[last] // "Вс"

	// Можно инициализировать массив указав индексы. Индексы в массиве выступают в качестве ключей
	days = [7]string{
		0: "Пн",
		1: "Вт",
		2: "Ср",
		3: "Чт",
		4: "Пт",
		5: "Сб",
		6: "Вс",
	}

	// Или только некоторые индексы
	days = [7]string{3: "Четверг", 1: "Вторник", 6: "Воскресенье"}
	fmt.Println(days) // ["" "Вторник" "" "Четверг" "" "" "Воскресенье"]

	// Можно сделать красивые синтаксические индексы
	const (
		Monday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	// Инициализация массива
	days = [...]string{
		Monday:    "Пн",
		Tuesday:   "Вт",
		Wednesday: "Ср",
		Thursday:  "Чт",
		Friday:    "Пт",
		Saturday:  "Сб",
		Sunday:    "Вс",
	}

	// Обращение
	monday = days[Monday] // "Пн"

	// Изменение
	days[Monday] = "Пн"

	// Итерация по массиву
	for i := 0; i < len(days); i++ {
		day := days[i]
		fmt.Println("Индекс:", i, "День:", day)
	}

	// Итерация с помощью range
	for i, day := range days {
		fmt.Println("Индекс:", i, "День:", day)
	}

	_, _, _, _ = length, capacity, monday, sunday
}
