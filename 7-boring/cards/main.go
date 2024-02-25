package main

import "fmt"

var deckSize int

func main() {
	// var card string = "Ace of Spades"
	// var card = "Ace of Spades"
	// card := "Ace of Spades"
	card := newCard()
	randomNumber := newRandomNumber()
	// deckSize = 50
	fmt.Println(card)
	fmt.Println(randomNumber)
	// fmt.Println(deckSize)
}

func newCard() string {
	return "Five of Diamonds"
}

func newRandomNumber() int {
	return 123
}
