package ranks

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	rank := fmt.Sprintf("%v", Jack)

	if rank != "J" {
		t.Errorf("Should be J")
	}
}

func TestFormatLong(t *testing.T) {
	rank := fmt.Sprintf("%v", Jack.Long())

	if rank != "Jack" {
		t.Errorf("Should be Jack")
	}
}

func TestCreateFourRank(t *testing.T) {
	rank, err := From("4")
	if err != nil {
		t.Errorf("Should be valid")
	}

	if rank.String() != "4" {
		t.Errorf("Should be 4")
	}
}

func TestCreateQueenRank(t *testing.T) {
	rank, err := From("Q")
	if err != nil {
		t.Errorf("Should be valid")
	}

	if rank.String() != "Q" {
		t.Errorf("Should be Q")
	}
}

func TestInvalidRank(t *testing.T) {
	_, err := From("M")
	if err == nil {
		t.Errorf("Should not be valid")
	}
}
