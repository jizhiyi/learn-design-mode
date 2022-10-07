package sample1

import (
	"learn/factorymethod/sample1/factory"
	"testing"
)

func TestIDCardFactory(t *testing.T) {
	idCardFactory := factory.NewIDCardFactory()
	card1 := idCardFactory.Create("小明")
	card2 := idCardFactory.Create("小红")
	card3 := idCardFactory.Create("小刚")
	card4 := idCardFactory.Create("小刚")
	card1.Use()
	card2.Use()
	card3.Use()
	card4.Use()
}
