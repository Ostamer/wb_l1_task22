package main

import (
	"fmt"
)

func main() {
	var a, b float64

	fmt.Println("Введите значение a (больше 2^20):")
	fmt.Scan(&a)
	fmt.Println("Введите значение b (больше 2^20):")
	fmt.Scan(&b)

	// Проверка условий
	if a <= (1<<20) || b <= (1<<20) {
		fmt.Println("Оба значения должны быть больше 2^20.")
		return
	}

	// Выполнение рассчетов
	summ := a + b
	minus := a - b
	multi := a * b
	var slash float64

	if b != 0 {
		slash = a / b
	} else {
		fmt.Println("Ошибка: Деление на ноль.")
		return
	}

	// Вывод результатов
	fmt.Printf("Сумма: %.2f\n", summ)
	fmt.Printf("Разность: %.2f\n", minus)
	fmt.Printf("Произведение: %.2f\n", multi)
	fmt.Printf("Частное: %.2f\n", slash)
}
