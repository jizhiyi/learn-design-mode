package sample1

import (
	"testing"
)

func TestPrintBanner(t *testing.T) {
	var p Print
	p = NewPrintBanner("Hello World")
	p.PrintWeak()
	p.PrintStrong()
}
