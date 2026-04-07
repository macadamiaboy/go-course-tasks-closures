// Задание 2: Аккумулятор
//
// Замыкание может не только ЧИТАТЬ захваченную переменную, но и ИЗМЕНЯТЬ её.
// Переменная живёт пока жива функция-замыкание - между вызовами значение сохраняется.
//
// Напиши функцию makeAdder() func(int) int, которая:
//   - хранит внутри переменную sum = 0
//   - каждый вызов прибавляет переданное число к sum
//   - возвращает новое значение sum
//
// Напиши функцию makeAdderWithReset() (func(int) int, func()), которая:
//   - делает то же самое
//   - НО дополнительно возвращает функцию reset(), которая обнуляет sum
//
// Ожидаемый вывод:
//   Первый аккумулятор:
//   +10 -> 10
//   +5  -> 15
//   +3  -> 18
//
//   Второй аккумулятор (независимый от первого):
//   +100 -> 100
//
//   Первый продолжает с 18:
//   +1 -> 19
//
//   Аккумулятор с ресетом:
//   +7 -> 7
//   +3 -> 10
//   reset!
//   +5 -> 5
//
// Запусти: go run main.go

package main

import "fmt"

// TODO: напиши функцию makeAdder() func(int) int
// Подсказка: объяви sum := 0 внутри makeAdder,
// и верни функцию которая меняет sum и возвращает его
func makeAdder() func(int) int {
	sum := 0
	return func(n int) int {
		sum += n
		return sum
	}
}

// TODO: напиши функцию makeAdderWithReset() (func(int) int, func())
// Подсказка: та же идея, но верни два значения - add и reset.
// Обе функции захватывают одну и ту же переменную sum.
func makeAdderWithReset() (func(int) int, func()) {
	sum := 0
	return func(n int) int {
			sum += n
			return sum
		}, func() {
			sum = 0
		}
}

func main() {
	// TODO: создай два независимых аккумулятора через makeAdder()
	// и проверь что они не мешают друг другу
	first := makeAdder()
	second := makeAdder()

	fmt.Println(first(6))
	fmt.Println(second(9))
	fmt.Println(first(10))
	fmt.Println(second(19))

	// TODO: создай аккумулятор с ресетом и проверь reset
	third, reset := makeAdderWithReset()
	fmt.Println(third(11))
	fmt.Println(third(12))
	reset()
	fmt.Println(third(10))
}
