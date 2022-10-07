package sample2

import "testing"

func TestAbstractDisplayDisplay(t *testing.T) {
	var charDisplay AbstractDisplayIn = NewCharDisplay('H')
	charDisplay.Display()
	var stringDisplay AbstractDisplayIn = NewStringDisplay("hello world")
	stringDisplay.Display()
}
