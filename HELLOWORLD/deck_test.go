package main

import (
	"os"
	"testing")


func TestNewDeck(t *testing.T ){

	d := newDeck()

	if len(d)!=52{

		t.Errorf("Expected length of 52 but got %v", len(d))

	}

	if d[0] !="Ace of Spades" {

		t.Errorf("Expected first card was Ace of Spade but got %v", d[0])

	}

	if d[len(d)-1] !="King of Clubs" {

		t.Errorf("Expected last card was King of Clubs but got %v", d[len(d)-1])

	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T){

	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck)!=52 {
		t.Errorf("Deck not loaded properly")
	}

	os.Remove("_decktesting")

}