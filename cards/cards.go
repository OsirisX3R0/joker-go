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

func CardFrom(rankAbbr string, suitAbbr string) (Card, error) {
	rank, err := ranks.RankFrom(rankAbbr)

	if err != nil {
		return Card{
			rank: ranks.DEFUALT,
			suit: suits.DEFAULT,
		}, err
	}

	suit, err := suits.SuitFrom(suitAbbr)

	if err != nil {
		return Card{
			rank: ranks.DEFUALT,
			suit: suits.DEFAULT,
		}, err
	}

	return Card{
		rank: rank,
		suit: suit,
	}, nil
}

func (c Card) String() string {
	return fmt.Sprintf("%v of %v", c.rank, c.suit)
}
