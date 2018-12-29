package blackjack

type Symbol string
type FaceCard int
type Sort string
type Card struct {
	Number int
	Symbol Symbol
}
type Deck struct {
	Set []Card
}

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
	c := NewCard(1, Heart)
	cs = append(cs, *c)
	d := &Deck{
		cs,
	}
	return d
}
