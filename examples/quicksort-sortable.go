// Based on: https://github.com/lfourky/go-quicksort/blob/master/quicksort.go
// package sort
package main

import (
	"fmt"
	"strings"
)

func main() {
	sortVampires()
	sortIntegers()
	sortStrings()
}

func sortIntegers() {
	list := ListOfIntegers{123, 789, 234, 987, 345, 890, 567, 678, 456, 876, 765, 654, 543, 432, 321}
	QuickSort(list, true)
	fmt.Println(list)
	QuickSort(list, false)
	fmt.Println(list)
}

type ListOfIntegers []int

func (l ListOfIntegers) LessEqual(i, j int) bool {
	return l[i] <= l[j]
}
func (l ListOfIntegers) Len() int {
	return len(l)
}
func (l ListOfIntegers) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
	return
}

func sortStrings() {
	list := ListOfStrings{
		"PL/1", "Fortran", "Assembler IBM/360", "COBOL", "C", "ML/1", "C++",
		"Pascal", "Modula-2", "Oberon", "cmd", "sh", "SQL", "FoxPro",
		"Perl-5", "JavaScript", "Java", "Ruby", "Python", "Kotlin", "Go",
	}
	QuickSort(list, true)
	fmt.Printf("%s\n", list)
	QuickSort(list, false)
	fmt.Printf("%s\n", list)
}

type ListOfStrings []string

func (l ListOfStrings) LessEqual(i, j int) bool {
	return strings.ToUpper(l[i]) <= strings.ToUpper(l[j])
}
func (l ListOfStrings) Len() int {
	return len(l)
}
func (l ListOfStrings) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
	return
}

// https://github.com/lfourky/go-quicksort/blob/master/quicksort.go
type vampires []vampire
type vampire struct {
	name         string
	prefersWomen bool
	killCount    int
}

func (v vampires) LessEqual(i, j int) bool { // Chooose your sorting key (wisely)
	return v[i].killCount <= v[j].killCount
}
func (v vampires) Len() int {
	return len(v)
}
func (v vampires) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func sortVampires() {
	vampz := make(vampires, 0, 2)
	vampz = append(vampz,
		vampire{name: "Dracula", prefersWomen: true, killCount: 100001},
		vampire{name: "Sava Savanović", prefersWomen: false, killCount: 3},
	)
	QuickSort(vampz, false)
	fmt.Println(vampz)
}

type sortable interface {
	LessEqual(i, j int) bool
	Len() int
	Swap(i, j int)
}

// QuickSort sorts using the QuickSort algorithm... who would've guessed :)
func QuickSort(s sortable, isAscending bool) {
	sort(s, 0, s.Len()-1, isAscending)
}

func sort(s sortable, start, end int, isAscending bool) {
	if start >= end {
		return
	}
	pivot := partition(s, start, end, isAscending)
	sort(s, start, pivot-1, isAscending)
	sort(s, pivot+1, end, isAscending)
}

func partition(s sortable, start, end int, isAscending bool) int {
	pivot := end
	partitionIndex := start
	for current := start; current < end; current++ {
		if (isAscending && s.LessEqual(current, pivot)) || (!isAscending && !s.LessEqual(current, pivot)) {
			s.Swap(current, partitionIndex)
			partitionIndex++
		}
	}
	s.Swap(partitionIndex, pivot)
	return partitionIndex
}
