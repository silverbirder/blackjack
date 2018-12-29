package blackjack

import (
	"testing"
)

func TestAceIs1(t *testing.T) {
	c := NewCard(1, 1, Heart, "1")
	actual := c.Score()
	expected := 1
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
