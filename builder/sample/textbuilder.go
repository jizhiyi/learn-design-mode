package sample

import (
	"fmt"
	"strings"
)

type TextBuilder struct {
	content *strings.Builder
}

func NewTextBuilder() *TextBuilder {
	return &TextBuilder{content: &strings.Builder{}}
}

func (text *TextBuilder) MakeTitle(title string) {
	fmt.Fprintln(text.content, "==============================")
	fmt.Fprintf(text.content, "『%s』\n\n", title)
}

func (text *TextBuilder) MakeString(str string) {
	fmt.Fprintf(text.content, "■%s\n", str)
}

func (text *TextBuilder) MakeItems(items []string) {
	for _, item := range items {
		fmt.Fprintf(text.content, "　・%s\n", item)
	}
	fmt.Fprintln(text.content, "")
}

func (text *TextBuilder) Close() {
	fmt.Fprintln(text.content, "==============================")
}

func (text *TextBuilder) GetResult() string {
	return text.content.String()
}
