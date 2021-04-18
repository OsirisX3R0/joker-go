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
	deck := NewBlankDeck()
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

	if deck.ExactEqual(otherDeck) != true {
		t.Errorf("Should be equal")
	}
}

func TestNewStandardDeckWithJokers(t *testing.T) {
	deck, err := NewStandardDeckWithJokers()
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

	joker, err := cards.From("X")
	if err != nil {
		t.Errorf("Should encounter no errors")
	}
	otherDeck.Push(joker)
	otherDeck.Push(joker)

	if deck.ExactEqual(otherDeck) != true {
		t.Errorf("Should be equal")
	}
}

func TestShuffle(t *testing.T) {
	standDeck, err := NewStandardDeck()
	if err != nil {
		t.Errorf("Should not encounter errors")
	}
	shuffleDeck, err := NewStandardDeck()
	if err != nil {
		t.Errorf("Should not encounter errors")
	}
	shuffleDeck.Shuffle()

	if standDeck.Equal(shuffleDeck) != true {
		t.Errorf("Should be equal")
	}
}

func TestShuffleExact(t *testing.T) {
	standDeck, err := NewStandardDeck()
	if err != nil {
		t.Errorf("Should not encounter errors")
	}
	shuffleDeck, err := NewStandardDeck()
	if err != nil {
		t.Errorf("Should not encounter errors")
	}
	shuffleDeck.Shuffle()

	if standDeck.ExactEqual(shuffleDeck) == true {
		t.Errorf("Should not be equal")
	}
}
