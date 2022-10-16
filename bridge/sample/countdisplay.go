package sample

type CountDisplay struct {
	*Display
}

func NewCountDisplay(impl DisplayImpl) *CountDisplay {
	return &CountDisplay{Display: NewDisplay(impl)}
}

func (c *CountDisplay) MultiDisplay(times int) {
	c.Open()
	for i := 0; i < times; i++ {
		c.Print()
	}
	c.Close()
}
