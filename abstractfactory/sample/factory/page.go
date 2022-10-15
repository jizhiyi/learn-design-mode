package factory

import (
	"bufio"
	"os"
)

type Page struct {
	Title  string
	Author string

	Contents []Item
	inner    Item
}

func NewPage(title string, author string, inner Item) *Page {
	return &Page{Title: title, Author: author, inner: inner}
}

func (p *Page) Add(item Item) {
	p.Contents = append(p.Contents, item)
}

func (p *Page) Output() error {
	f, err := os.Create(p.Title + ".html")
	if err != nil {
		return err
	}
	bufWriter := bufio.NewWriter(f)
	_, err = bufWriter.WriteString(p.MakeHTML())
	if err != nil {
		return err
	}
	err = bufWriter.Flush()
	if err != nil {
		return err
	}
	err = f.Close()
	return err
}

func (p *Page) MakeHTML() string {
	if p.inner != nil {
		return p.inner.MakeHTML()
	}
	panic("not over")
}
