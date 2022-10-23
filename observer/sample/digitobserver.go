package sample

import (
	"fmt"
	"time"
)

type DigitObserver struct {
}

func (d *DigitObserver) Updata(n NumberGenerator) {
	fmt.Println("DigitObserver:", n.GetNumber())
	time.Sleep(100 * time.Millisecond)
}
