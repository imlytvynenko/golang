package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
// In OO terms it's smth like 'extends'
type deck []string

// Here we define 'receiver'
// Any variable of type 'deck' now gets access to the 'print' method
// In OO tems 'receiver' is very close to 'this' of 'self' keyword
// By convencion the name of receiver starts from first/first+second letters of type name
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}