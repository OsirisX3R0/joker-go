package decks

import "github.com/OsirisX3R0/joker-go/cards"

type Pile []cards.Card

type IPile interface {
	Push(cards.Card)
	Pull() cards.Card
	Size() int
	Contains(cards.Card) bool
	ExactEqual(Pile) bool
	Equal(Pile) bool
}

func (p Pile) Push(c cards.Card) {
	p = append(p, c)
}

func (p Pile) Pull() cards.Card {
	card, discard := p[0], p[1:]
	p = discard

	return card
}

func (p Pile) Size() int {
	return len(p)
}

func (p Pile) Contains(c cards.Card) bool {
	for i := range p {
		if p[i].Equal(c) {
			return true
		}
	}

	return false
}

func (p Pile) Equal(o Pile) bool {
	equal := true

	if ((p == nil) != (o == nil)) ||
		p.Size() != o.Size() {
		equal = false
	}

	for i := range p {
		if !equal {
			continue
		}

		if !p.Contains(o[i]) {
			equal = false
		}
	}

	return equal
}

func (p Pile) ExactEqual(o Pile) bool {
	equal := true

	if ((p == nil) != (o == nil)) ||
		p.Size() != o.Size() {
		equal = false
	}

	for i := range p {
		if !equal {
			continue
		}

		if p[i] != o[i] {
			equal = false
		}
	}

	return equal
}
