package blackjack

import (
	"testing"
)

func TestDeckSortDescAsc(t *testing.T) {
	d := NewDeck()
	d.Sort(Desc)
	if d.Set[0].Score() < d.Set[len(SymbolList)+1].Score() {
		t.Errorf("not desc sort first:%v, second:%v", d.Set[0], d.Set[len(SymbolList)+1])
	}
	d.Sort(Asc)
	if d.Set[0].Score() > d.Set[len(SymbolList)+1].Score() {
		t.Errorf("not asc sort first:%v, second:%v", d.Set[0], d.Set[len(SymbolList)+1])
	}
}
