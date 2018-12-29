package blackjack

type Symbol string

const (
	Heart   Symbol = "Heart"
	Club    Symbol = "Club"
	Spade   Symbol = "Spade"
	Diamond Symbol = "Diamond"
)

type Card struct {
	Number      int
	Value       int
	Symbol      Symbol
	DisplayText string
}

func NewCard(number int, value int, symbol Symbol, displayText string) *Card {
	c := &Card{Number: number, Value: value, Symbol: symbol, DisplayText: displayText}
	return c
}

func (c *Card) Score() int {
	return c.Value
}
