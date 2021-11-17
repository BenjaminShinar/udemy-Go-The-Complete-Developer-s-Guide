package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	l := len(d)
	exp := 16
	if l != exp {
		t.Errorf("Expected length %d, got %v", exp, l)
	}
	// first := d[0]
	// last := d[l-1]
	expFirst, expLast := "Ace of Spades", "Four of Diamonds"
	first, last := d[0], d[l-1]
	if first != expFirst {
		t.Errorf("expected first card to be %v, was %v", expFirst, first)
	}
	if last != expLast {
		t.Errorf("expected last card to be %v, was %v", expLast, last)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
    filename := "_deckTesting"
    os.Remove(filename)
	d := newDeck()
	
    d.saveToFile(filename)
	loadedDeck := newDeckFromFile(filename)

	l1 := len(d)
	l2 := len(loadedDeck)
	if l1 != l2 {
		t.Errorf("Expected length %d, got %v", l1,l2)
	}
    os.Remove(filename)
}