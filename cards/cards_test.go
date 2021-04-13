package cards

import (
	"fmt"
	"testing"
)

func TestCreateAceOfSpades(t *testing.T) {
	var card, err = From("A", "S")
	if err != nil {
		t.Errorf("Should encounter no errors")
	}

	if fmt.Sprintf("%v", card) != "AS" {
		t.Errorf("Should be Ace of Spades (AS)")
	}
}

func TestInvalidCard(t *testing.T) {
	var _, err = From("X", "X")
	if err == nil {
		t.Errorf("Should encounter an error")
	}
}
