package sample

import (
	"learn/facade/sample/pagemaker"
	"testing"
)

func Test_PageMaker(t *testing.T) {
	err := pagemaker.MakeWelcomePage("jizhihong@hyuki.com", "welcome.html")
	if err != nil {
		t.Fatal(err)
	}
}
