package main

import "go-test-sample/blackjack"

func main() {
	u := blackjack.NewUser("hoge", false)
	d := blackjack.NewDeck()
	u.Draw(d.Set, true)
}
