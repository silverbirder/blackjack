package blackjack

type Algorithm interface {
	Think(user User) bool
}

type SimpleAlgorithm struct{}

func (s *SimpleAlgorithm) Think(user User) bool {
	score := user.TotalScore()
	return score < BustScore-5
}
