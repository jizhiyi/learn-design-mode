package sample

import "testing"

func Test_BigString(t *testing.T) {
	bs := NewBigString("1234567")
	bs.Print()
}
