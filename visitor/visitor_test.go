package visitor

import (
	"fmt"
	"learn/visitor/sample"
	"testing"
)

func Test_Visitor(t *testing.T) {
	rootdir := sample.NewDirectory("root")
	bindir := sample.NewDirectory("bin")
	tmpdir := sample.NewDirectory("tmp")
	usrdir := sample.NewDirectory("usr")
	rootdir.Add(bindir)
	rootdir.Add(tmpdir)
	rootdir.Add(usrdir)
	bindir.Add(sample.NewFile("vi", 10000))
	bindir.Add(sample.NewFile("latex", 10000))
	rootdir.Accept(&sample.ListVisitor{})

	fmt.Println()
	fmt.Println("Making user entries...")
	yuki := sample.NewDirectory("yuki")
	hanako := sample.NewDirectory("hanako")
	tomura := sample.NewDirectory("tomura")
	usrdir.Add(yuki)
	usrdir.Add(hanako)
	usrdir.Add(tomura)
	yuki.Add(sample.NewFile("diary.html", 100))
	yuki.Add(sample.NewFile("Composite.java", 200))
	hanako.Add(sample.NewFile("memo.tex", 300))
	tomura.Add(sample.NewFile("game.doc", 400))
	tomura.Add(sample.NewFile("junk.mail", 500))
	rootdir.Accept(&sample.ListVisitor{})
}
