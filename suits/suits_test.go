package suits

import (
	"fmt"
	"testing"
)

func TestCreateClubsSuit(t *testing.T) {
	suit, err := From("C")

	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	if suit != Clubs {
		t.Errorf("Should be Clubs")
	}
}

func TestCreateSpadesSuit(t *testing.T) {
	suit, err := From("S")

	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	if suit != Spades {
		t.Errorf("Should be Spades")
	}
}

func TestCreateHeartsSuit(t *testing.T) {
	suit, err := From("H")

	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	if suit != Hearts {
		t.Errorf("Should be Hearts")
	}
}

func TestCreateDiamondsSuit(t *testing.T) {
	suit, err := From("D")

	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	if suit != Diamonds {
		t.Errorf("Should be Diamonds")
	}
}

func TestCreateJokerSuit(t *testing.T) {
	suit, err := From("")

	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	if suit != Joker {
		t.Errorf("Should be Joker")
	}
}

func TestInvalidSuit(t *testing.T) {
	_, err := From("X")

	if err == nil {
		t.Errorf("Should encounter error")
	}

}

func TestFormat(t *testing.T) {
	rank := fmt.Sprintf("%v", Hearts)

	if rank != "H" {
		t.Errorf("Should be H")
	}
}

func TestFormatLong(t *testing.T) {
	rank := fmt.Sprintf("%v", Hearts.Long())

	if rank != "Hearts" {
		t.Errorf("Should be Hearts")
	}
}
