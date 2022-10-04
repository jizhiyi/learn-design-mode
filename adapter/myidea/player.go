package myidea

type Player struct {
	Name  string
	Skill string
	Coins int
}

func NewPlayer(name string, skill string, coins int) *Player {
	return &Player{Name: name, Skill: skill, Coins: coins}
}

func (p *Player) PlayerDoSkill() {
	DoSkill(p)
}

func (p *Player) PlayerDoPay() {
	DoPay(p)
}
