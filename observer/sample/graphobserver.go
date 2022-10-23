package sample

import (
	"fmt"
	"time"
)

type GraphObserver struct {
}

func (g *GraphObserver) Updata(n NumberGenerator) {
	fmt.Print("GraphObserver:")
	count := n.GetNumber()
	for i := 0; i < count; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	time.Sleep(100 * time.Millisecond)
}
