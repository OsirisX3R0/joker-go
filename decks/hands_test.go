package decks

import (
	"testing"
)

func TestNewHand(t *testing.T) {
	hand := NewHand()
	var otherHand Hand

	if !hand.Equal(otherHand) {
		t.Errorf("Should be equal")
	}
}
