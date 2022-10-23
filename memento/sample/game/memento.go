package game

type Memento struct {
	money  int
	fruits []string
}

func (m *Memento) setMoney(money int) {
	m.money = money
}

func (m *Memento) GetMoney() int {
	return m.money
}

func (m *Memento) addFruit(fruit string) {
	m.fruits = append(m.fruits, fruit)
}

func (m *Memento) getFruit() []string {
	return m.fruits
}
