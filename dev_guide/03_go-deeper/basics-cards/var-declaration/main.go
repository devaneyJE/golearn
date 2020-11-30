package main

import "fmt"

func main() {
	var card string = "Ace of Spades" // extended format, very specific in typing
	cardTwo := "King of Hearts"
	cardTwo = "Queen of Hearts" // only use : when declaring/initializing; don't need when reassigning
	cardThree := newCard()      // can declare and assign var with function

	fmt.Println(card, cardTwo, cardThree)
}

func newCard() string { // need to declare return type of function
	return "Jack of Diamonds"
}
