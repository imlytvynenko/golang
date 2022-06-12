package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
// In OO terms it's smth like 'extends'
type deck []string

func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Here we define 'receiver'
// Any variable of type 'deck' now gets access to the 'print' method
// In OO tems 'receiver' is very close to 'this' of 'self' keyword
// By convencion the name of receiver starts from first/first+second letters of type name
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Slice Range syntax: sliceName[startIndex:notIncludedEndIndex]
// Return multiple values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
