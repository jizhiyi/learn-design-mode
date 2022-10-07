package sample2

import "testing"

func TestAbstractDisplayDisplay(t *testing.T) {
	var charDisplay AbstractDisplayIn
	charDisplay = NewCharDisplay('H')
	charDisplay.Display()
	var stringDisplay AbstractDisplayIn
	stringDisplay = NewStringDisplay("hello world")
	stringDisplay.Display()
}
