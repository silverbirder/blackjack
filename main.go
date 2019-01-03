package main

import (
	"blackjack/blackjack"
)

func main() {
	users := make([]blackjack.User, 0)
	me := blackjack.NewUser("me", false)
	you := blackjack.NewUser("you", true)
	you.SetAlgorithm(&blackjack.SimpleAlgorithm{})
	users = append(users, *me)
	users = append(users, *you)
	deck := blackjack.NewDeck()
	deck.Sort(blackjack.Random)

	g := blackjack.NewGame(users, *deck)
	g.InitTurn()
	g.MainTurn()
	g.JudgeTurn()
}
