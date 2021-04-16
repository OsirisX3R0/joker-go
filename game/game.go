package game

import "github.com/OsirisX3R0/joker-go/decks"

type Game struct {
	deck decks.Deck
}

type GameOptions struct {
	numberOfDecks int
	jokers        bool
}

type BaseGame interface {
	Generate(GameOptions)
}

func (g Game) Generate(o GameOptions) {
	// do nothing by default
	return
}
