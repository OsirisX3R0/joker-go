package game

import (
	"fmt"

	"github.com/OsirisX3R0/joker-go/cards"
	"github.com/OsirisX3R0/joker-go/decks"
	"github.com/OsirisX3R0/joker-go/ranks"
	"github.com/OsirisX3R0/joker-go/suits"
)

type DefaultGame struct {
	Game
}

func (g DefaultGame) Generate(o GameOptions) error {
	deck := decks.New()
	for i := 1; i <= o.numberOfDecks; i++ {
		for _, rank := range ranks.RankEnum() {
			rankStr := fmt.Sprintf("%v", rank)
			for _, suit := range suits.SuitEnum() {
				suitStr := fmt.Sprintf("%v", suit)
				card, err := cards.From(rankStr, suitStr)
				if err != nil {
					return err
				}
				deck.Push(card)
			}
		}
	}

	return nil
}
