package blackjack

type Game struct {
	Players []User
	Deck    Deck
}

func NewGame(users []User, deck Deck) *Game {
	g := &Game{
		Players: users,
		Deck:    deck,
	}
	return g
}

func (g *Game) InitTurn() {
	for _, p := range g.Players {
		g.Deck.Set = p.Draw(g.Deck.Set, true)
	}
}
