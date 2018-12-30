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
}

func NewUser(name string, auto bool) *User {
	c := make([]Card, 0)
	u := &User{
		Name:  name,
		Hands: c,
		Auto:  auto,
	}
	return u
}

func (u *User) Draw(Cards []Card, auto bool) []Card {
	if auto || u.Auto || isDraw() {
		pop := Cards[len(Cards)-1]
		u.Hands = append(u.Hands, pop)
		return Cards[:len(Cards)-1]
	}
	return Cards
}

func isDraw() bool {
	for {
		ans := readLine("do you draw cards? y/n")
		if ans == "y" {
			fmt.Println("ok. start draw")
			return true
		} else if ans == "n" {
			fmt.Println("ok. stop draw")
			return false
		} else {
			fmt.Println("please, input y or n")
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
