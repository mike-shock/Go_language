package main

import "fmt"

type BaseType = string

func main() {
	var v BaseType = "Значение" // переменная типа BaseType
	var p *BaseType             // указатель на переменную типа BaseType

	p = &v  // ссылка на значение переменной типа BaseType
	c := *p // значение переменной типа BaseType по ссылке на v

	fmt.Println(v, p, *p, c)

	changeValue(p)

	fmt.Println(v, p, *p, c)
}

func changeValue(p *BaseType) {
	var newValue BaseType = "Новое значение"
	*p = newValue
}
