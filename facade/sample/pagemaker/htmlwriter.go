package pagemaker

import (
	"bufio"
	"io"
	"os"
)

type htmlWriter struct {
	writer *bufio.Writer
	io.Closer
}

func NewHtmlWriter(fileName string) (*htmlWriter, error) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	hw := &htmlWriter{}
	hw.Closer = f
	hw.writer = bufio.NewWriter(f)
	return hw, nil
}

func (hw *htmlWriter) title(title string) {
	hw.writer.WriteString("<html>")
	hw.writer.WriteString("<head>")
	hw.writer.WriteString("<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" />")
	hw.writer.WriteString("<title>" + title + "</title>")
	hw.writer.WriteString("</head>")
	hw.writer.WriteString("<body>\n")
	hw.writer.WriteString("<h1>" + title + "</h1>\n")
}

func (hw *htmlWriter) paragraph(msg string) {
	hw.writer.WriteString("<p>" + msg + "</p>\n")
}

func (hw *htmlWriter) link(href, caption string) {
	hw.paragraph("<a href=\"" + href + "\">" + caption + "</a>")
}

func (hw *htmlWriter) mailto(mailaddr, username string) {
	hw.link("mailto:"+mailaddr, username)
}
func (hw *htmlWriter) close() {
	hw.writer.WriteString("</body>")
	hw.writer.WriteString("</html>\n")
	hw.writer.Flush()
	hw.Close()
}
