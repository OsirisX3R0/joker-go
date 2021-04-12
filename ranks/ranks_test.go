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
