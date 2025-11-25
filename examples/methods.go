package main

import "fmt"

type Celsius float32

func (t Celsius) String() string { return fmt.Sprintf("%g°C", t) }

func main() {
	var t Celsius = 37.0
	println(t.String()) // 37°C

	recorder := TapeRecorder{Model: "Нота 203-1 стерео"}
	recorder.play(a)
}

type Album struct {
	name, artist string
	year, length int
	media        string
}

var a = Album{"Dark Side of the Moon", "Pink Floyd", 1973, 44, "катушка 18 см"}

type TapeRecorder struct {
	Model string
}

func (r TapeRecorder) play(a Album) {
	fmt.Printf("Playing album '%s' by '%s' for %d minutes...\n",
		a.name, a.artist, a.length)
}
