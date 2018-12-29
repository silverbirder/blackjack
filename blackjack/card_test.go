package blackjack

import (
	"testing"
)

func TestAceIs1(t *testing.T) {
	c := NewCard(1, Heart)
	actual := c.Score()
	expected := 1
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestFaceCardIs10(t *testing.T) {
	for _, faceCard := range FaceCardList {
		c := NewCard(int(faceCard), Heart)
		actual := c.Score()
		expected := 10
		if actual != expected {
			t.Errorf("actual %v\nwant %v", actual, expected)
		}
	}
}

func TestDeckSortDesc(t *testing.T) {
	//d := NewDeck()
	//d.Sort(Desc)
	//if d.Set[0].Score() > d.Set[1].Score() {
	//	t.Errorf("not desc sort %v > %v", d.Set[0], d.Set[1])
	//}
}
