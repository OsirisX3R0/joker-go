package suits

import "errors"

type Suit int

const (
	Clubs Suit = iota
	Spades
	Hearts
	Diamonds
	Joker
)

func SuitFrom(abbr string) (Suit, error) {
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
		return Joker, errors.New("Not a valid suit")
	}
}

func (s Suit) String() string {
	switch s {
	case Clubs:
		return "C"
	case Spades:
		return "S"
	case Hearts:
		return "H"
	case Diamonds:
		return "D"
	default:
		return ""
	}
}

func (s Suit) Long() string {
	switch s {
	case Clubs:
		return "Clubs"
	case Spades:
		return "Spades"
	case Hearts:
		return "Hearts"
	case Diamonds:
		return "Diamonds"
	default:
		return ""
	}
}
