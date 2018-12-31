package blackjack

import (
	"testing"
)

func TestGame_InitGame(t *testing.T) {
	users := make([]User, 0)
	me := NewUser("me", true)
	you := NewUser("you", true)
	users = append(users, *me)
	users = append(users, *you)
	deck := NewDeck()
	deck.Sort(Asc)
	originalDecLen := len(deck.Set)
	g := NewGame(users, *deck)
	g.InitTurn()
	actual := len(g.Deck.Set)
	expected := originalDecLen - 4
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestGame_IsAllPlayerDraw(t *testing.T) {
	users := make([]User, 0)
	me := NewUser("me", true)
	you := NewUser("you", true)
	users = append(users, *me)
	users = append(users, *you)
	deck := NewDeck()
	deck.Sort(Asc)
	g := NewGame(users, *deck)
	g.Deck.Set, g.Players[0].Hands, g.Players[0].End = g.Players[0].Draw(g.Deck.Set, false)
	g.Deck.Set, g.Players[0].Hands, g.Players[0].End = g.Players[0].Draw(g.Deck.Set, false)
	g.Deck.Set, g.Players[0].Hands, g.Players[0].End = g.Players[0].Draw(g.Deck.Set, false)
	g.Deck.Set, g.Players[1].Hands, g.Players[1].End = g.Players[1].Draw(g.Deck.Set, false)
	g.Deck.Set, g.Players[1].Hands, g.Players[1].End = g.Players[1].Draw(g.Deck.Set, false)
	g.Deck.Set, g.Players[1].Hands, g.Players[1].End = g.Players[1].Draw(g.Deck.Set, false)
	actual := g.IsAllPlayerDraw()
	expected := true
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
