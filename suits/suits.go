package suits

import (
	"errors"
	"fmt"
)

type Suit int

const (
	Clubs Suit = iota
	Spades
	Hearts
	Diamonds
	Joker
	INVALID
)

func From(abbr string) (Suit, error) {
	errorText := fmt.Sprintf("%v Not a valid suit", abbr)

	switch abbr {
	case "C":
		return Clubs, nil
	case "S":
		return Spades, nil
	case "H":
		return Hearts, nil
	case "D":
		return Diamonds, nil
	case "":
		return Joker, nil
	default:
		return INVALID, errors.New(errorText)
	}
}

func (s Suit) String() string {
	return [...]string{
		"C",
		"S",
		"H",
		"D",
	}[s]
}

func (s Suit) Long() string {
	return [...]string{
		"Clubs",
		"Spades",
		"Hearts",
		"Diamonds",
	}[s]
}
