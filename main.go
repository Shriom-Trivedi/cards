package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 3)

	fmt.Println("hand:", hand, "\nremainingDeck", remainingDeck)
}
