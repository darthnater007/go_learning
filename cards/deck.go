package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type, deck: a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// deal specified number of cards to return hand and remaining
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//shuffle the deck
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// print contents of deck type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// cast deck slice to single string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save deck to file
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// read from file to create new deck
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}
