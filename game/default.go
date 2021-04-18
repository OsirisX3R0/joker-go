package game

import (
	"github.com/OsirisX3R0/joker-go/decks"
)

type DefaultGame struct {
	Game
}

func (g DefaultGame) Generate(o GameOptions) {
	deck, err := decks.NewStandardDeckWithJokers()
	if err != nil {
		panic(err)
	}

	g.deck = deck
	hand := decks.NewHand()
	g.hands = append(g.hands, hand)
}
