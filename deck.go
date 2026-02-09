package main

import "fmt"

type deck []string

// () receiver
// any var of type deck gat access to the print method
// d is an actual copy, every var can call this function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
