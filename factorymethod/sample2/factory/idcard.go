package factory

import "fmt"

type IDCard struct {
	owner string
}

func newIDCard(owner string) *IDCard {
	fmt.Printf("制作 %s 的ID卡\n", owner)
	return &IDCard{owner: owner}
}

func (idCard *IDCard) Use() {
	fmt.Printf("使用 %s 的ID卡\n", idCard.owner)
}

func (idCard *IDCard) GetOwner() string {
	return idCard.owner
}
