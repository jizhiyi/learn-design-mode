package listfactory

import (
	"learn/abstractfactory/sample/factory"
	"strings"
)

type ListPage struct {
	*factory.Page
}

func (l *ListPage) MakeHTML() string {
	stringbuf := strings.Builder{}
	stringbuf.WriteString("<html><head><meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" /><title>" + l.Title + "</title></head>")
	stringbuf.WriteString("<body>\n")
	stringbuf.WriteString("<h1>" + l.Title + "</h1>\n")
	stringbuf.WriteString("<ul>\n")
	for _, content := range l.Contents {
		stringbuf.WriteString(content.MakeHTML())
	}
	stringbuf.WriteString("</ul>\n")
	stringbuf.WriteString("<hr><address>" + l.Author + "</address>")
	stringbuf.WriteString("</body></html>\n")
	return stringbuf.String()
}
