package blackjack

import (
	"testing"
)

func TestUser_Draw(t *testing.T) {
	u := NewUser("me", true)
	d := NewDeck()
	d.Sort(Desc)
	originalCardsLen := len(d.Set)
	c, h, _ := u.Draw(d.Set, true)
	actual := len(h)
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

func TestUser_TotalScore(t *testing.T) {
	u := NewUser("me", true)
	d := NewDeck()
	d.Sort(Desc)
	d.Set, u.Hands, u.End = u.Draw(d.Set, true)
	d.Set, u.Hands, u.End = u.Draw(d.Set, true)
	actual := u.TotalScore()
	expected := 2
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestUser_IsBust(t *testing.T) {
	u := NewUser("me", true)
	d := NewDeck()
	d.Sort(Asc)
	d.Set, u.Hands, u.End = u.Draw(d.Set, true)
	d.Set, u.Hands, u.End = u.Draw(d.Set, true)
	d.Set, u.Hands, u.End = u.Draw(d.Set, true)
	actual := u.isBust(BustScore)
	expected := true
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
