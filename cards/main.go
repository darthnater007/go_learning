package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" //Longform variable declaration
	card := "Ace of Spades"   // use := to have go 'infer' the variable type
	card = "Five of Diamonds" // do not use := to reassign values

	fmt.Println(card)
}
