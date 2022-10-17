package sample

import "math/rand"

type WinningStrategy struct {
	r       *rand.Rand
	won     bool
	preHand *Hand
}

func NewWinningStrategy() *WinningStrategy {
	return &WinningStrategy{r: rand.New(rand.NewSource(123))}
}

func (w *WinningStrategy) NextHand() *Hand {
	if !w.won {
		w.preHand = NewHand(w.r.Intn(3))
	}
	return w.preHand
}

func (w *WinningStrategy) Study(win bool) {
	w.won = win
}
