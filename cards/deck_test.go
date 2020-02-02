package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	fmt.Println("\nNew Deck Creation:")

	if len(d) != 52 {
		t.Errorf("\tExpected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("\tExpected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("\tExpected last card to be King of Clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	saveDeck := newDeck()
	saveDeck.saveToFile("_decktesting")

	loadDeck := newDeckFromFile("_decktesting")

	if len(loadDeck) != 52 {
		t.Errorf("\tExpected 52 cards in deck, but got %v", len(loadDeck))
	}

	os.Remove("_decktesting")
}
