// https://ru.wikipedia.org/wiki/Ответ_на_главный_вопрос_жизни,_вселенной_и_всего_такого
package main

import (
	"fmt"
	"time"
)

const Вопрос = "Главный вопрос жизни, вселенной и вообще"
const Ответ = 42

func Мыслитель(вопрос string) int {
	return Ответ
}

func main() {
	var время_работы = 7_500_000

	начало := time.Now()
	окончание := начало.AddDate(время_работы, 0, 0)

	ответ := Мыслитель(Вопрос)
	fmt.Printf("Начав в %v и проработав %v лет, Мыслитель в %v дал  такой ответ: '%v'.\n",
                   начало, время_работы, окончание, ответ)
}
