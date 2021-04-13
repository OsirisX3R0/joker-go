package cards

import (
	"fmt"
	"testing"
)

func TestCreateAceOfSpades(t *testing.T) {
	card, err := From("A", "S")
	if err != nil {
		t.Errorf("Should encounter no errors")
	}

	if fmt.Sprintf("%v", card) != "AS" {
		t.Errorf("Should be Ace of Spades (AS)")
	}
}

func TestCreateJoker(t *testing.T) {
	card, err := From("X")
	if err != nil {
		t.Errorf("Should encounter no errors")
	}

	if fmt.Sprintf("%v", card) != "Jk" {
		t.Errorf("Should be Ace of Spades (Jk)")
	}
}

func TestInvalidCard(t *testing.T) {
	_, err := From("X", "X")
	if err == nil {
		t.Errorf("Should encounter an error")
	}
}

func TestFormat(t *testing.T) {
	card, err := From("Q", "H")
	if err != nil {
		t.Errorf("Should be valid")
	}

	cardStr := fmt.Sprintf("%v", card)

	if cardStr != "QH" {
		t.Errorf("Should be QH")
	}
}

func TestFormatLong(t *testing.T) {
	card, err := From("Q", "H")
	if err != nil {
		t.Errorf("Should be valid")
	}

	cardStr := fmt.Sprintf("%v", card.Long())

	if cardStr != "Queen of Hearts" {
		t.Errorf("Should be Queen of Hearts")
	}
}
