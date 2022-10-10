package sample

import (
	"testing"
)

func TestText_Builder(t *testing.T) {
	textBuilder := NewTextBuilder()
	director := NewDirector(textBuilder)
	director.Construct()
	t.Log(textBuilder.GetResult())
}

func TestHTML_Builder(t *testing.T) {
	htmlBuilder := NewHTMLBuilder()
	director := NewDirector(htmlBuilder)
	director.Construct()
	t.Log(htmlBuilder.GetResult())
}
