package sample

import (
	"fmt"
	"learn/prototype/sample/framework"
)

type UnderlinePen struct {
	ulchar byte
}

func NewUnderlinePen(ulchar byte) *UnderlinePen {
	return &UnderlinePen{ulchar: ulchar}
}

// 功能方法
func (u *UnderlinePen) Use(str string) {
	sz := len(str)
	fmt.Printf("\"%s\"\n ", str)
	for i := 0; i < sz; i++ {
		fmt.Printf("%c", u.ulchar)
	}
	fmt.Println()
}

// 复制方法
func (u *UnderlinePen) Clone() framework.Product {
	return &UnderlinePen{ulchar: u.ulchar}
}
