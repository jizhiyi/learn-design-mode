package sample

import "testing"

func TestDisplay(t *testing.T) {
	d1 := NewDisplay(NewStringDisplayImpl("Hello, China."))
	d2 := NewDisplay(NewStringDisplayImpl("Hello, World."))
	d3 := NewCountDisplay(NewStringDisplayImpl("Hello, Universe."))
	d1.DisplayFun()
	d2.DisplayFun()
	d3.DisplayFun()
	d3.MultiDisplay(5)
}
