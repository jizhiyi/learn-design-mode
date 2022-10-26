package sample

const (
	MoveType_Left  int = -1
	MoveType_Right int = 1
	MoveType_Stay  int = 0
)

type Move struct {
	moveType int
}

func NewMove(moveType int) *Move {
	return &Move{moveType: moveType}
}

func (m *Move) Execute(pos int) int {
	return pos + m.moveType
}
