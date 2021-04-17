package hands

import (
	"github.com/OsirisX3R0/joker-go/cards"
	"github.com/OsirisX3R0/joker-go/decks"
)

type Hand []cards.Card

func New() Hand {
	var hand Hand

	return hand
}

func (h Hand) Push(c cards.Card) {
	h = append(h, c)
}

func (h *Hand) Draw(d *decks.Deck) {
	card := d.Draw()

	h.Push(card)
}

func (h Hand) Pull() cards.Card {
	card, hand := h[0], h[1:]
	h = hand

	return card
}

func (h *Hand) Discard(d *decks.Deck) {
	card := h.Pull()

	d.Discard(card)
}
