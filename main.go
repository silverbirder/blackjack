package main

import (
	"fmt"
	"go-test-sample/blackjack"
)

func main() {
	c := blackjack.NewCard(1, blackjack.Heart)
	fmt.Println(c.Score())
}
