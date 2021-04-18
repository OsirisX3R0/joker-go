package game

import (
	"github.com/OsirisX3R0/joker-go/decks"
)

type Game struct {
	deck  decks.Deck
	hands []decks.Hand
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
	g.deck = decks.NewBlankDeck()
	hand := decks.NewHand()
	g.hands = append(g.hands, hand)
}
