package strategy

import (
	"learn/strategy/sample"
	"testing"
)

func Test_Strategy(t *testing.T) {
	player1 := sample.NewPlayer("jizhihong", sample.NewWinningStrategy())
	player2 := sample.NewPlayer("jizhiyi", sample.NewProbStrategy())
	for i := 0; i < 1000; i++ {
		hand1 := player1.NextHand()
		hand2 := player2.NextHand()
		if hand1.IsStrongerThan(hand2) {
			t.Log("Winer: ", player1)
			player1.Win()
			player2.Lose()
		} else if hand2.IsStrongerThan(hand1) {
			t.Log("Winer: ", player2)
			player1.Lose()
			player2.Win()
		} else {
			t.Log("Even...")
			player1.Even()
			player2.Even()
		}
	}
	t.Log("Total result:")
	t.Log(player1)
	t.Log(player2)
}
