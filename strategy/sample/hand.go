package sample

const (
	HANDVALUE_GUU int = 0
	HANDVALUE_CHO int = 1
	HANDVALUE_PAA int = 2
)

var staticHand []*Hand = []*Hand{
	{handvalue: HANDVALUE_GUU},
	{handvalue: HANDVALUE_CHO},
	{handvalue: HANDVALUE_PAA},
}

type Hand struct {
	handvalue int
}

func NewHand(handvalue int) *Hand {
	return staticHand[handvalue]
}

func (self *Hand) IsStrongerThan(h *Hand) bool {
	return self.fight(h) == 1
}

func (self *Hand) fight(h *Hand) int {
	if self.handvalue == h.handvalue {
		return 0
	} else if (self.handvalue + 1) == h.handvalue {
		return 1
	}
	return -1
}
