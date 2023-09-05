package main

// import "fmt"

func main() {

	cards := newDeckFromFile("Test.txt")
	cards.shuffle()
	cards.print()
	//hand, remainingDeck := deal(cards, 5)
	// hand.print()

	// remainingDeck.print()
	// cards.saveToFile("Test.txt")

}
