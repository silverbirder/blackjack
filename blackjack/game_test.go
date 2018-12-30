package blackjack

import (
	"testing"
)

func TestGame_GameInitTurn(t *testing.T) {
	users := make([]User, 0)
	me := NewUser("me", true)
	you := NewUser("you", true)
	users = append(users, *me)
	users = append(users, *you)
	deck := NewDeck()
	deck.Sort(Desc)
	g := NewGame(users, *deck)
	g.InitTurn()
}
