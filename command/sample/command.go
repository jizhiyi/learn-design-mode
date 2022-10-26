package sample

type Command interface {
	Execute(pos int) int
}
