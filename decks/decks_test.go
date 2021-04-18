package decks

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/OsirisX3R0/joker-go/cards"
	"github.com/OsirisX3R0/joker-go/ranks"
	"github.com/OsirisX3R0/joker-go/suits"
)

func TestNewDeck(t *testing.T) {
	deck := New()
	var blankPile Pile
	otherDeck := Deck{
		draw:    blankPile,
		discard: blankPile,
	}

	if reflect.DeepEqual(deck, otherDeck) != true {
		t.Errorf("Should be equal")
	}
}

func TestNewStandardDeck(t *testing.T) {
	deck, err := NewStandardDeck()
	if err != nil {
		t.Errorf("Should encounter no errors")
	}

	var otherDeck Deck
	for _, rank := range ranks.RankEnum() {
		rankStr := fmt.Sprintf("%v", rank)
		for _, suit := range suits.SuitEnum() {
			suitStr := fmt.Sprintf("%v", suit)
			card, err := cards.From(rankStr, suitStr)
			if err != nil {
				t.Errorf("Should encounter no errors")
			}
			otherDeck.Push(card)
		}
	}

	if deck.Equal(otherDeck) != true {
		t.Errorf("Should be equal")
	}
}
