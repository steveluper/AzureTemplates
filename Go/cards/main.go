package main

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards,5)
	// hand.print()
	// remainingCards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// Card from File Demo
	// cards := newDeckFromFile("my_cards")
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}