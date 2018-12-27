package main

import (
	"./blackjack"
)

func main() {
	c := blackjack.NewCard(1, 1, blackjack.Heart, "1")
	c.Score()
}
