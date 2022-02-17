package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")

	savedCards := newDeckFromFile("my_cards")
	savedCards.print()

	savedCards.shuffle()
	savedCards.print()
}
