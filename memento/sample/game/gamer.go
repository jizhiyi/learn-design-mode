package game

import (
	"fmt"
	"math/rand"
	"strings"
)

var fruitsName []string = []string{
	"苹果", "葡萄", "香蕉", "橘子",
}

type Gamer struct {
	money  int
	fruits []string
}

func NewGamer(money int) *Gamer {
	return &Gamer{money: money}
}

func (g *Gamer) GetMoney() int {
	return g.money
}

func (g *Gamer) Bet() {
	dice := rand.Intn(6) + 1
	if dice == 1 {
		g.money += 100
		fmt.Println("所持金钱增加了。")
	} else if dice == 2 {
		g.money /= 2
		fmt.Println("所持金钱减半了。")
	} else if dice == 6 {
		f := getFruit()
		fmt.Printf("获得了水果(%s)。\n", f)
		g.fruits = append(g.fruits, f)
	} else {
		fmt.Println("什么都没有发生。")
	}
}

func (g *Gamer) CreateMemento() *Memento {
	m := &Memento{}
	m.setMoney(g.money)
	for _, fruit := range g.fruits {
		if strings.HasPrefix(fruit, "好吃的") {
			m.addFruit(fruit)
		}
	}
	return m
}

func (g *Gamer) RestoreMemento(memento *Memento) {
	g.money = memento.GetMoney()
	g.fruits = memento.getFruit()
}

func (g *Gamer) String() string {
	return fmt.Sprintf("[money = %d, fruits = %v]", g.money, g.fruits)
}

func getFruit() string {
	var prefix string
	if rand.Intn(2) == 1 {
		prefix = "好吃的"
	}
	return prefix + fruitsName[rand.Intn(len(fruitsName))]
}
