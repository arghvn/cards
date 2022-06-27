package main

func main() {
	// former  commands
	// cards := newDeck()
	// cards.savetofile("my_cards")
	// create a new deck of cards
	cards := newDeckfromfile("my_cards")
	cards.Print()
	// refer to shuffle
}
