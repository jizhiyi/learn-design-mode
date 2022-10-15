package factory

type Tray struct {
	ItemSt
	Trays []Item
	inner Item
}

func NewTray(caption string, inner Item) *Tray {
	return &Tray{ItemSt: ItemSt{Caption: caption}, inner: inner}
}

func (t *Tray) Add(item Item) {
	t.Trays = append(t.Trays, item)
}

func (t *Tray) MakeHTML() string {
	if t.inner != nil {
		return t.inner.MakeHTML()
	}
	panic("not implemented") // TODO: Implement
}
