package factory

type Link struct {
	ItemSt
	Url   string
	inner Item
}

func NewLink(caption string, url string, inner Item) *Link {
	return &Link{ItemSt: ItemSt{Caption: caption}, Url: url, inner: inner}
}

func (l *Link) MakeHTML() string {
	if l.inner != nil {
		return l.inner.MakeHTML()
	}
	panic("not over")
}
