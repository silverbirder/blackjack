package blackjack

import (
	"testing"
)

func TestCard_AceIs1(t *testing.T) {
	c := NewCard(1, Heart)
	actual := c.Score()
	expected := 1
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestCard_FaceCardIs10(t *testing.T) {
	for _, faceCard := range FaceCardList {
		c := NewCard(int(faceCard), Heart)
		actual := c.Score()
		expected := 10
		if actual != expected {
			t.Errorf("actual %v\nwant %v", actual, expected)
		}
	}
}
