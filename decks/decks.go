package decks

import "github.com/OsirisX3R0/joker-go/cards"

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
