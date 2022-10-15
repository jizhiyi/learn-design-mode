package factory

type Factory interface {
	CreateLink(caption string, url string) *Link
	CreateTray(caption string) *Tray
	CreatePage(title string, author string) *Page
}
