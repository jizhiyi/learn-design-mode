package framework

type Manager struct {
	products map[string]Product
}

func NewManager() *Manager {
	return &Manager{products: make(map[string]Product)}
}

func (m *Manager) Register(protoname string, p Product) {
	m.products[protoname] = p
}

func (m *Manager) Create(protoname string) Product {
	p, ok := m.products[protoname]
	if ok {
		return p.Clone()
	}
	return nil
}
