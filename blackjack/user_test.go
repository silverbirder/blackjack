package blackjack

import (
	"testing"
)

func TestUser_Draw(t *testing.T) {
	u := NewUser("me", true)
	d := NewDeck()
	d.Sort(Desc)
	originalCardsLen := len(d.Set)
	c := u.Draw(d.Set, true)
	actual := len(u.Hands)
	expected := 1
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
	actual = len(c)
	expected = originalCardsLen - 1
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
