package main

// import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	// cards.saveToFile("My_cards")
}
