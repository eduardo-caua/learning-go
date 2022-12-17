package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf(`Expected deck length of 16, but got %v`, len(d))
	}

	if d[0] != "Ace of Spaces" {
		t.Errorf(`Expected first carf of Ace of Spaces, but got %v`, d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf(`Expected laster card of Four of Clubs, but got %v`, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	const filename string = "_decktesting"

	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	ld := newDeckFromFile(filename)

	if len(ld) != 16 {
		t.Errorf(`Expected deck length of 16, but got %v`, len(ld))
	}

	os.Remove(filename)
}
