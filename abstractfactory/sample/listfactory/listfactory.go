package listfactory

import "learn/abstractfactory/sample/factory"

type ListFactory struct {
}

func (l *ListFactory) CreateLink(caption string, url string) *factory.Link {
	listLink := &ListLink{}
	listLink.Link = factory.NewLink(caption, url, listLink)
	return listLink.Link
}

func (l *ListFactory) CreateTray(caption string) *factory.Tray {
	liskTray := &ListTray{}
	liskTray.Tray = factory.NewTray(caption, liskTray)
	return liskTray.Tray
}

func (l *ListFactory) CreatePage(title string, author string) *factory.Page {
	listPage := &ListPage{}
	listPage.Page = factory.NewPage(title, author, listPage)
	return listPage.Page
}
