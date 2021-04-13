package cards

import (
	"fmt"

	"github.com/OsirisX3R0/joker-go/ranks"
	"github.com/OsirisX3R0/joker-go/suits"
)

type Card struct {
	rank ranks.Rank
	suit suits.Suit
}

var InvalidCard Card = Card{
	rank: ranks.INVALID,
	suit: suits.INVALID,
}

func From(rankAbbr string, suitAbbr string) (Card, error) {
	rank, err := ranks.From(rankAbbr)

	if err != nil {
		return InvalidCard, err
	}

	suit, err := suits.From(suitAbbr)

	if err != nil {
		return InvalidCard, err
	}

	return Card{
		rank: rank,
		suit: suit,
	}, nil
}

func (c Card) String() string {
	return fmt.Sprintf("%v%v", c.rank, c.suit)
}

func (c Card) Long() string {
	return fmt.Sprintf("%v of %v", c.rank.Long(), c.suit.Long())
}
