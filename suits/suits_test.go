package suits

import "testing"

func TestCreateClubsSuit(t *testing.T) {
	suit, err := SuitFrom("C")

	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	if suit != Clubs {
		t.Errorf("Should be Clubs")
	}
}

func TestCreateSpadesSuit(t *testing.T) {
	suit, err := SuitFrom("S")

	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	if suit != Spades {
		t.Errorf("Should be Spades")
	}
}

func TestCreateHeartsSuit(t *testing.T) {
	suit, err := SuitFrom("H")

	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	if suit != Hearts {
		t.Errorf("Should be Hearts")
	}
}

func TestCreateDiamondsSuit(t *testing.T) {
	suit, err := SuitFrom("D")

	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	if suit != Diamonds {
		t.Errorf("Should be Diamonds")
	}
}

func TestCreateJokerSuit(t *testing.T) {
	suit, err := SuitFrom("")

	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	if suit != Joker {
		t.Errorf("Should be Joker")
	}
}

func TestInvalidSuit(t *testing.T) {
	_, err := SuitFrom("X")

	if err == nil {
		t.Errorf("Should encounter error")
	}

}
