package decks

import (
	"fmt"

	"github.com/OsirisX3R0/joker-go/cards"
	"github.com/OsirisX3R0/joker-go/ranks"
	"github.com/OsirisX3R0/joker-go/suits"
)

type Deck struct {
	draw    Pile
	discard Pile
}

type Pile []cards.Card

func New() Deck {
	var blankPile Pile

	return Deck{
		draw:    blankPile,
		discard: blankPile,
	}
}

func (d *Deck) Push(c cards.Card) {
	d.draw = append(d.draw, c)
}

func (d *Deck) Draw() cards.Card {
	card, draw := d.draw[0], d.draw[1:]
	d.draw = draw

	return card
}

func (d *Deck) Pull() cards.Card {
	card, discard := d.discard[0], d.discard[1:]
	d.discard = discard

	return card
}

func (d *Deck) Discard(c cards.Card) {
	d.discard = append(d.discard, c)
}

func (d Deck) DrawSize() int {
	return len(d.draw)
}

func (d Deck) DiscardSize() int {
	return len(d.discard)
}

func (d Deck) Equal(o Deck) bool {
	equal := true

	if ((d.draw == nil) != (o.draw == nil)) ||
		d.DrawSize() != o.DrawSize() ||
		((d.discard == nil) != (o.discard == nil)) ||
		d.DiscardSize() != o.DiscardSize() {
		equal = false
	}

	for i := range d.draw {
		if equal {
			continue
		}

		if d.draw[i] != o.draw[i] {
			equal = false
		}
	}

	for i := range d.discard {
		if equal {
			continue
		}

		if d.discard[i] != o.discard[i] {
			equal = false
		}
	}

	return equal
}

func (d Deck) String() string {
	deckStr := ""
	for _, card := range d.draw {
		cardStr := fmt.Sprintf("%v\n", card)
		deckStr += cardStr
	}

	return deckStr
}

func (d Deck) Verbose() string {
	deckStr := ""
	for _, card := range d.draw {
		cardStr := fmt.Sprintf("%v\n", card.Long())
		deckStr += cardStr
	}

	return deckStr
}

func NewStandardDeck() (Deck, error) {
	deck := New()
	// for i := 1; i <= o.numberOfDecks; i++ {
	for _, rank := range ranks.RankEnum() {
		rankStr := fmt.Sprintf("%v", rank)
		for _, suit := range suits.SuitEnum() {
			suitStr := fmt.Sprintf("%v", suit)
			card, err := cards.From(rankStr, suitStr)
			if err != nil {
				return deck, err
			}
			deck.Push(card)
		}
	}
	// }

	return deck, nil
}

func NewStandardDeckWithJokers() (Deck, error) {
	deck, err := NewStandardDeck()
	if err != nil {
		return deck, err
	}

	joker, err := cards.From("X")
	if err != nil {
		return deck, err
	}

	deck.Push(joker)
	deck.Push(joker)

	return deck, nil
}
