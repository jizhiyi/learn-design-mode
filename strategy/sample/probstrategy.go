package sample

import "math/rand"

type ProbStrategy struct {
	r                *rand.Rand
	prevHandValue    int
	currentHandValue int
	history          [][]int
}

func NewProbStrategy() *ProbStrategy {
	p := &ProbStrategy{}
	p.r = rand.New(rand.NewSource(2))
	p.history = [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	return p
}

func (p *ProbStrategy) NextHand() *Hand {
	bet := p.r.Intn(p.getSum(p.currentHandValue))
	handvalue := 0
	if bet < p.history[p.currentHandValue][0] {
		handvalue = 0
	} else if bet < p.history[p.currentHandValue][0]+p.history[p.currentHandValue][1] {
		handvalue = 1
	} else {
		handvalue = 2
	}
	p.prevHandValue = p.currentHandValue
	p.currentHandValue = handvalue
	return NewHand(handvalue)
}

func (p *ProbStrategy) getSum(hv int) int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += p.history[hv][i]
	}
	return sum
}

func (p *ProbStrategy) Study(win bool) {
	if win {
		p.history[p.prevHandValue][p.currentHandValue]++
	} else {
		p.history[p.prevHandValue][(p.currentHandValue+1)%3]++
		p.history[p.prevHandValue][(p.currentHandValue+2)%3]++
	}
}
