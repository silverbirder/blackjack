package blackjack

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name  string
	Hands []Card
	Auto  bool
	End   bool
}

func NewUser(name string, auto bool) *User {
	c := make([]Card, 0)
	u := &User{
		Name:  name,
		Hands: c,
		Auto:  auto,
		End:   false,
	}
	return u
}

func (u *User) Draw(Cards []Card, auto bool) ([]Card, []Card, bool) {
	if auto || u.Auto {
		pop := Cards[len(Cards)-1]
		u.Hands = append(u.Hands, pop)
		return Cards[:len(Cards)-1], u.Hands, u.End
	}
	if u.End = !isDraw(); !u.End {
		pop := Cards[len(Cards)-1]
		u.Hands = append(u.Hands, pop)
		return Cards[:len(Cards)-1], u.Hands, u.End
	}
	return Cards, u.Hands, u.End
}

func isDraw() bool {
	for {
		ans := readLine("do you draw cards? y/n .")
		if ans == "y" {
			fmt.Println("ok. start draw.")
			return true
		} else if ans == "n" {
			fmt.Println("ok. stop draw.")
			return false
		} else {
			fmt.Println("please, input y or n.")
		}
	}
}

func readLine(msg string) string {
	fmt.Println(msg)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		ans := s.Text()
		return ans
	}
	if s.Err() != nil {
		// non-EOF error.
		log.Fatal(s.Err())
	}
	return ""
}

func (u *User) String() string {
	return u.Name
}

func (u *User) TotalScore() int {
	score := 0
	for _, c := range u.Hands {
		score = score + c.Score()
	}
	return score
}

func (u *User) isBust(bustScore int) bool {
	return u.TotalScore() > bustScore
}
