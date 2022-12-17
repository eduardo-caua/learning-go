package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	d1, d2 := deal(cards, 9)
	d1.print()
	d2.print()
}
