package sample

type Strategy interface {
	NextHand() *Hand
	Study(win bool)
}
