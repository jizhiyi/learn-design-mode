package sample

import "math/rand"

type RandomNumberGenerator struct {
	*NumberGeneratorBase
	rand   rand.Rand
	number int
}

func NewRandomNumberGenerator() *RandomNumberGenerator {
	r := &RandomNumberGenerator{}
	r.NumberGeneratorBase = &NumberGeneratorBase{
		NumberGenerator: r,
	}
	return r
}

func (r *RandomNumberGenerator) GetNumber() int {
	return r.number
}

func (r *RandomNumberGenerator) Execute() {
	for i := 0; i < 20; i++ {
		r.number = rand.Intn(50)
		r.NotifyObserver()
	}
}
