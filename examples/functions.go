package main

import "fmt"

type F1 func(int, int) int // тип функции = её сигнатура

// у любой функции есть тип, например: func(int, int) int
func add(x, y int) int { return x + y } // соответствует типу F1

func returnsFunc() F1 {
	return add
}

func receivesFunc(a, b int, f F1) (r int) { r = f(a, b); return r }

func main() {

	// функция как значение переменной
	var f1 F1 = add // присваивание объявленной функции
	fmt.Println(f1(40, 2))

	var fa = func() { println("anonymous") } // анонимная функция типа func()
	fa()

	// функция как возвращаемое значение
	f2 := returnsFunc()
	fmt.Println(f2(40, 2))

	// функция как параметр
	sum := receivesFunc(21, 21, add)
	production := receivesFunc(21, 2, func(x, y int) int { return x * y })
	fmt.Println(sum, production)

	// определение и вызов анонимной функции
	func() { println("lambda") }() // lambda типа func()
}
