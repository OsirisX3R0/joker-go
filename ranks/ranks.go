package ranks

import (
	"errors"
	"fmt"
)

type Rank int

const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
	Joker
	INVALID
)

func RankFrom(abbr string) (Rank, error) {
	errorText := fmt.Sprintf("%v not a valid Rank", abbr)

	switch abbr {
	case "2":
		return Two, nil
	case "3":
		return Three, nil
	case "4":
		return Four, nil
	case "5":
		return Five, nil
	case "6":
		return Six, nil
	case "7":
		return Seven, nil
	case "8":
		return Eight, nil
	case "9":
		return Nine, nil
	case "10":
		return Ten, nil
	case "J":
		return Jack, nil
	case "Q":
		return Queen, nil
	case "K":
		return King, nil
	case "A":
		return Ace, nil
	case "":
		return Joker, nil
	default:
		return INVALID, errors.New(errorText)
	}
}

func (r Rank) String() string {
	return [...]string{
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"J",
		"Q",
		"K",
		"A",
		"J",
	}[r]
}

func (r Rank) Long() string {
	return [...]string{
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
		"Ace",
		"Joker",
	}[r]
}
