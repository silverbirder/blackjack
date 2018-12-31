package blackjack

import (
	"testing"
)

func TestAlgorithm_Think(t *testing.T) {
	u := NewUser("me", true)
	d := NewDeck()
	d.Sort(Asc)
	d.Set, u.Hands, _, _ = u.Draw(d.Set, true)
	d.Set, u.Hands, _, _ = u.Draw(d.Set, true)
	s := &SimpleAlgorithm{}
	isThink := s.Think(*u)
	actual := isThink
	expected := false
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
