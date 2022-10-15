package listfactory

import (
	"fmt"
	"learn/abstractfactory/sample/factory"
)

type ListLink struct {
	*factory.Link
}

func (l *ListLink) MakeHTML() string {
	return fmt.Sprintf("  <li><a href=\"%s\">%s</a></li>\n", l.Url, l.Caption)
}
