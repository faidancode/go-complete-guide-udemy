package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 32 {
		t.Errorf("Expected deck length of 32, but got %v", len(d))
	}
}
