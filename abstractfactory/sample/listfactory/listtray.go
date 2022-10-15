package listfactory

import (
	"learn/abstractfactory/sample/factory"
	"strings"
)

type ListTray struct {
	*factory.Tray
}

func (l *ListTray) MakeHTML() string {
	stringbuf := strings.Builder{}
	stringbuf.WriteString("<li>\n")
	stringbuf.WriteString(l.Caption + "\n")
	stringbuf.WriteString("<ul>\n")
	for _, item := range l.Trays {
		stringbuf.WriteString(item.MakeHTML())
	}
	stringbuf.WriteString("</ul>")
	stringbuf.WriteString("</li>")
	return stringbuf.String()
}
