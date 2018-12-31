package blackjack

import "fmt"

type Game struct {
	Players []User
	Deck    Deck
}

const BustScore = 21

func NewGame(users []User, deck Deck) *Game {
	g := &Game{
		Players: users,
		Deck:    deck,
	}
	return g
}

func (g *Game) InitTurn() {
	fmt.Println("start blackjack.")
	fmt.Println("init turn.")
	for index, p := range g.Players {
		for i := 0; i < 2; i++ {
			g.Deck.Set, g.Players[index].Hands, g.Players[index].End = p.Draw(g.Deck.Set, true)
			// hide second card if user is auto mode.
			if p.Auto && i == 1 {
				fmt.Println(p.String() + " draw ???(hidden).")
			} else {
				fmt.Println(p.String() + " draw " + p.Hands[len(p.Hands)-1].String() + ".")
			}
		}
	}
}

func (g *Game) MainTurn() {
	for index, p := range g.Players {
		for {
			g.Deck.Set, g.Players[index].Hands, g.Players[index].End = p.Draw(g.Deck.Set, false)
			if p.End {
				break
			}
			if p.isBust(BustScore) {
				fmt.Println(p.String() + " bust.")
				break
			}
			fmt.Println(p.String() + " draw " + p.Hands[len(p.Hands)-1].String() + ".")
		}
	}
}
func (g *Game) JudgeTurn() {
	isGameDraw := g.IsAllPlayerDraw()
	if isGameDraw {
		fmt.Println("this game is draw.")
		return
	}
	min := BustScore
	var winUser User
	for _, p := range g.Players {
		pMin := BustScore - p.TotalScore()
		if pMin < 0 {
			break
		}
		if min > pMin {
			winUser = p
		}
	}
	fmt.Printf("%v (score:%v) win.", winUser.String(), winUser.TotalScore())
}

func (g *Game) IsAllPlayerDraw() bool {
	for _, p := range g.Players {
		if !p.isBust(BustScore) {
			return false
		}
	}
	return true
}
