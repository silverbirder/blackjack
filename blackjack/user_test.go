package blackjack

import (
	"testing"
)

func TestUser_Draw(t *testing.T) {
	u := NewUser("me")
	d := NewDeck()
	d.Sort(Desc)
	originalCardsLen := len(d.Set)
	c := u.Draw(d.Set)
	actual := len(u.hands)
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
