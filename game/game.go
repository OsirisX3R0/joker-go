package game

import "github.com/OsirisX3R0/joker-go/decks"

type Game struct {
	deck decks.Deck
}

type BaseGame interface {
	Generate()
}

func (g Game) Generate() {
	// do nothing by default
	return
}
