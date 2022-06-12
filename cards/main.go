package main

func main() {
	cards := newDeck()

	// Assign multiple values from method that returns 2 values
	hand, ramainingDeck := deal(cards, 5)
	hand.print()
	ramainingDeck.print()
}
