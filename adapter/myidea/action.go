package myidea

import "fmt"

func DoSkill(p *Player) {
	fmt.Printf("player:%s, doSkill: %s\n", p.Name, p.Skill)
}

func DoPay(p *Player) {
	fmt.Printf("player:%s, pay %d coins\n", p.Name, p.Coins)
}
