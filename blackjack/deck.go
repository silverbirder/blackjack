package blackjack

import (
	"math/rand"
	"sort"
	"time"
)

type Deck struct {
	Set Cards
}
type Cards []Card
type Sort string

const (
	Random Sort = "Random"
	Desc   Sort = "Desc"
	Asc    Sort = "Asc"
)

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
			return d.Set[i].Score() < d.Set[j].Score()
		})
	case Random:
		rand.Seed(time.Now().UnixNano())
		for i := len(d.Set) - 1; i >= 0; i-- {
			j := rand.Intn(i + 1)
			d.Set[i], d.Set[j] = d.Set[j], d.Set[i]
		}
	}
}
