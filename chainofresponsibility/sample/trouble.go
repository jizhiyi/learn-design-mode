package sample

import "fmt"

type Trouble struct {
	number int
}

func (t *Trouble) GetNumber() int {
	return t.number
}

func (t *Trouble) String() string {
	return fmt.Sprintf("[Trouble: %d]", t.number)
}
