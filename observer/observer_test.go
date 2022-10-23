package observer

import (
	"learn/observer/sample"
	"testing"
)

func Test_Observer(t *testing.T) {
	generator := sample.NewRandomNumberGenerator()
	observer1 := &sample.DigitObserver{}
	observer2 := &sample.GraphObserver{}
	generator.AddObserver(observer1)
	generator.AddObserver(observer2)
	generator.Execute()
}
