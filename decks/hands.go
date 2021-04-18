package decks

import (
	"github.com/OsirisX3R0/joker-go/cards"
)

type Hand Pile

func NewHand() Hand {
	var hand Hand

	return hand
}

func (h Hand) Push(c cards.Card) {
	h = append(h, c)
}

func (h *Hand) Draw(d *Deck) {
	card := d.Draw()

	h.Push(card)
}

func (h Hand) Pull() cards.Card {
	card, hand := h[0], h[1:]
	h = hand

	return card
}

func (h *Hand) Discard(d *Deck) {
	card := h.Pull()

	d.Discard(card)
}

func (h Hand) Size() int {
	return len(h)
}

func (h Hand) Contains(c cards.Card) bool {
	for i := range h {
		if h[i].Equal(c) {
			return true
		}
	}

	return false
}

func (h Hand) Equal(o Hand) bool {
	equal := true

	if ((h == nil) != (o == nil)) ||
		h.Size() != o.Size() {
		equal = false
	}

	for i := range h {
		if !equal {
			continue
		}

		if !h.Contains(o[i]) {
			equal = false
		}
	}

	return equal
}

func (h Hand) ExactEqual(o Hand) bool {
	equal := true

	if ((h == nil) != (o == nil)) ||
		h.Size() != o.Size() {
		equal = false
	}

	for i := range h {
		if !equal {
			continue
		}

		if h[i] != o[i] {
			equal = false
		}
	}

	return equal
}
