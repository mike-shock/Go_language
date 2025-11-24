package main

import "fmt"

func sendMessage(ch chan<- string, s string) {
	ch <- s
}

func main() {
	var club chan string     // объявить канал для строк, значение nil
	club = make(chan string) // выделить память каналу для строк

	// club <- "Разговор о Go"		// будет deadlock !!!
	// отправить в канал значение	из одновременно выполняемой функции
	go func() { club <- "Предложен разговор о языке Go" }()
	received := <-club // получить значение из канала в переменную
	fmt.Println(received)

	go sendMessage(club, "Разговор о языке Go запланирован.")
	go sendMessage(club, "Разговор о языке Go состоялся.")

	m2, m1 := <-club, <-club // получить 2 сообщения
	close(club)
	message, ok := <-club // проверить доступность канала
	if !ok {              // канал закрыт
		fmt.Println("Разговор завершился.")
	}

	fmt.Println(m1, m2, message, ok)
}
