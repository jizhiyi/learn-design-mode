package sample1

import "testing"

func TestAbstractDisplay(t *testing.T) {
	d1 := NewCharDisplay('H')
	d2 := NewStringDisplay("hello world")
	d1.Display()
	d2.Display()
}
