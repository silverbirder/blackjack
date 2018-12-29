package blackjack

type User struct {
	Name  string
	hands []Card
}

func NewUser(name string) *User {
	c := make([]Card, 0)
	u := &User{
		Name:  name,
		hands: c,
	}
	return u
}

func (u *User) Draw(Cards []Card) []Card {
	pop := Cards[len(Cards)-1]
	u.hands = append(u.hands, pop)
	return Cards[:len(Cards)-1]
}
