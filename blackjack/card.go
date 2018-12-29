package blackjack

import (
	"math/rand"
	"sort"
)

type Card struct {
	Number int
	Symbol Symbol
}
type Deck struct {
	Set Cards
}
type Cards []Card
type Symbol string
type FaceCard int
type Sort string

const (
	Heart   Symbol = "Heart"
	Club    Symbol = "Club"
	Spade   Symbol = "Spade"
	Diamond Symbol = "Diamond"
)
const (
	Jack  FaceCard = 11
	Queen FaceCard = 12
	King  FaceCard = 13
)
const (
	Random Sort = "Random"
	Desc   Sort = "Desc"
	Asc    Sort = "Asc"
)
const FaceCardValue = 10

var SymbolList = []Symbol{
	Heart,
	Club,
	Spade,
	Diamond,
}
var FaceCardList = []FaceCard{
	Jack,
	Queen,
	King,
}

func NewCard(number int, symbol Symbol) *Card {
	c := &Card{
		Number: number,
		Symbol: symbol,
	}
	return c
}
func (c *Card) Score() int {
	if c.Number == int(Jack) ||
		c.Number == int(Queen) ||
		c.Number == int(King) {
		return FaceCardValue
	}
	return c.Number
}
func NewDeck() *Deck {
	cs := make([]Card, 0)
	for number := 1; number <= 13; number++ {
		for _, symbol := range SymbolList {
			c := NewCard(number, symbol)
			cs = append(cs, *c)
		}
	}
	d := &Deck{
		cs,
	}
	return d
}

func (d *Deck) Sort(sortType Sort) {
	switch sortType {
	case Desc:
		sort.SliceStable(d.Set, func(i, j int) bool {
			return d.Set[i].Score() >= d.Set[j].Score()
		})
	case Asc:
		sort.SliceStable(d.Set, func(i, j int) bool {
			return d.Set[i].Score() <= d.Set[j].Score()
		})
	case Random:
		for i := len(d.Set) - 1; i >= 0; i-- {
			j := rand.Intn(i + 1)
			d.Set[i], d.Set[j] = d.Set[j], d.Set[i]
		}
	}
}
