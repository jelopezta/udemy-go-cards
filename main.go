package main

import "fmt"

func main() {
	cards := newDeck()

	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	savedCards := newDeckFromFile("yeah")
	fmt.Println(savedCards)
}
