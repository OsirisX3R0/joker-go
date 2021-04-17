package game

import (
	"github.com/OsirisX3R0/joker-go/decks"
	"github.com/OsirisX3R0/joker-go/hands"
)

type Game struct {
	deck  decks.Deck
	hands []hands.Hand
}

type GameOptions struct {
	numberOfDecks int
	jokers        bool
}

type BaseGame interface {
	Generate(GameOptions)
}

func (g Game) Generate(o GameOptions) {
	// create blank properties by default default
	g.deck = decks.New()
	hand := hands.New()
	g.hands = append(g.hands, hand)
}
