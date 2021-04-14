package cards

import (
	"errors"
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

func From(abbrs ...string) (Card, error) {
	if len(abbrs) > 2 {
		return InvalidCard, errors.New("cards.From accepts no more than 2 arguments")
	}

	rankAbbr := abbrs[0]
	var suitAbbr string

	if len(abbrs) > 1 {
		suitAbbr = abbrs[1]
	} else {
		suitAbbr = ""
	}

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
	if c.rank == ranks.Joker {
		return "Jk"
	}
	return fmt.Sprintf("%v%v", c.rank, c.suit)
}

func (c Card) Long() string {
	if c.rank == ranks.Joker {
		return "Joker"
	}
	return fmt.Sprintf("%v of %v", c.rank.Long(), c.suit.Long())
}
