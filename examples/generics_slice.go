// Based on: https://hyperskill.org/learn/step/23505
package main

import (
	"errors"
	"fmt"
)

var EmptySliceError = errors.New("Slice is empty!")

func First[T any](a []T) (result T, err error) {
	if len(a) == 0 {
		return result, EmptySliceError
	}
	return a[0], err
}

func Last[T any](a []T) (result T, err error) {
	if len(a) == 0 {
		return result, EmptySliceError
	}
	return a[len(a)-1], err
}

func main() {
	sliceOfIntegers := []int{1, 2, 3, 4, 5}
	firstI, err := First[int](sliceOfIntegers)
	lastI, err := Last(sliceOfIntegers)
	fmt.Println(firstI, lastI, err)

	sliceOfStrings := []string{"Вышел", "зайчик", "погулять"}
	fmt.Println(Last[string](sliceOfStrings))
	fmt.Println(Last(sliceOfStrings))

	fmt.Println(First([]float32{}))
	fmt.Println(Last([]string{}))

	firstB, err := First([]bool{})
	lastB, err := First([]bool{})
	fmt.Println(firstB, lastB, err)
}
