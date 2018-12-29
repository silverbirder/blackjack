package blackjack

type Card struct {
	Number int
	Symbol Symbol
}
type Symbol string
type FaceCard int

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
