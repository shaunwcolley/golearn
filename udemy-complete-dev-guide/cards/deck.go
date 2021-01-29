package main

import "fmt"

// Create a new type of deck
// Slice of string
type deck []string

// a receiver in parentheses
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
