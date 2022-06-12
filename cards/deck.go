package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	} 

	ss := strings.Split(string(bs), ",")

	return deck(ss)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// one line swap
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
