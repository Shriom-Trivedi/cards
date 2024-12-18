package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %s", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
    os.Remove("_decktesting")

    deck := newDeck()
    deck.saveToFile("_decktesting")

    loadedDeck := newDeckFromFile("_decktesting")
    
    if len(loadedDeck) != 52 {
        t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
    }

    os.Remove("_decktesting")
}
