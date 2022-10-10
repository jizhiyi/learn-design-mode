package sample

import (
	"io"
	"os"
)

type HTMLBuilder struct {
	filename string
	wc       io.WriteCloser
}

func NewHTMLBuilder() *HTMLBuilder {
	return &HTMLBuilder{}
}

func (html *HTMLBuilder) MakeTitle(title string) {
	html.filename = title + ".html"
	f, err := os.Create(html.filename)
	if err != nil {
		return
	}
	html.wc = f
	html.wc.Write([]byte("<html><head><meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" /><title>" + title + "</title></head><body>"))
	html.wc.Write([]byte("<h1>" + title + "</h1>"))
}

func (html *HTMLBuilder) MakeString(str string) {
	html.wc.Write([]byte("<p>" + str + "</p>"))
}

func (html *HTMLBuilder) MakeItems(items []string) {
	html.wc.Write([]byte("<ul>"))
	for _, item := range items {
		html.wc.Write([]byte("<li>" + item + "</li>"))
	}
	html.wc.Write([]byte("</ul>"))
}

func (html *HTMLBuilder) Close() {
	html.wc.Close()
}

func (html *HTMLBuilder) GetResult() string {
	return html.filename
}
