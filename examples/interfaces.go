package main

import "fmt"

type Flyer interface{ fly() string } // 1-й интерфейс с методом
// все типы, которые реализуют метод fly(), будут соответствовать типу Flyer

type Bird struct{ Name string } // пользовательский тип Bird
func (b Bird) fly() string { // реализует метод fly() и поэтому
	return "flying…" // соответствует интерфейсу Flyer
}

type Swimmer interface{ swim() string } // 2-й интерфейс с методом

type Penguin struct{ Name string } // пользовательский тип Penguin
func (f Penguin) swim() string     { return "swimming..." }            // соответствует сразу
func (b Penguin) fly() string      { return "I can fly under water!" } // двум интерфейсам

func main() {
	var s = Bird{"Sparrow"}
	var p = Penguin{"Gentoo"}

	birds := []Flyer{s, p, Bird{"Dove"}} // Flyer — это тип данных
	for _, b := range birds {            // polymorphism
		fmt.Println(b, b.fly())
	}
}
