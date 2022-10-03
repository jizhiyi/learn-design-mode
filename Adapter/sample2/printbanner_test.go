package sample2

import (
	"testing"
)

func TestPrintBanner(t *testing.T) {
	// 这样的写法是有问题的,因为golang的继承好像不是很好使得样子
	p := *NewPrintBanner("Hello World")
	p.PrintWeak()
	p.PrintStrong()
}
