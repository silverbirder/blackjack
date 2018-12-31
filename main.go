package main

import "go-test-sample/blackjack"

func main() {
	users := make([]blackjack.User, 0)
	me := blackjack.NewUser("me", false)
	you := blackjack.NewUser("you", true)
	users = append(users, *me)
	users = append(users, *you)
	deck := blackjack.NewDeck()
	deck.Sort(blackjack.Desc)

	g := blackjack.NewGame(users, *deck)
	g.InitTurn()
	g.MainTurn()
	g.JudgeTurn()
}
