package main

import "time"

func f(n int) { println(n) }

func main() {
	for n := range 5 {
		go f(n + 1) // запустить 5 экземпляров f() одновременно c main()

		// можно запустить анонимные функции
		//		go func() { println(n + 1) }()

	}

	time.Sleep(5 * time.Second)
	println("Вышел зайчик погулять.")
}
