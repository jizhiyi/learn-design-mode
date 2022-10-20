package chainofresponsibility

import (
	"learn/chainofresponsibility/sample"
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
	alice := sample.NewNoSupport("Alice")
	bob := sample.NewLimitSupport("Bob", 100)
	charlie := sample.NewSpecialSupport("Charlie", 429)
	diana := sample.NewLimitSupport("Diana", 200)
	elmo := sample.NewOddSupport("Elmo")
	fred := sample.NewLimitSupport("Fred", 300)

	alice.SetNext(bob).SetNext(charlie).SetNext(diana).SetNext(elmo).SetNext(fred)

	for i := 0; i < 1000; i += 33 {
		alice.SupportFun(sample.NewTrouble(i))
	}
}
