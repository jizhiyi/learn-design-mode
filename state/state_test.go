package state

import (
	"learn/state/sample"
	"testing"
	"time"
)

func Test_State(t *testing.T) {
	worker := &sample.Worker{}
	for {
		for i := 0; i < 24; i++ {
			worker.SetHour(i)
			worker.DoThing()
			time.Sleep(500 * time.Millisecond)
		}
	}
}
