package proxy

import (
	"learn/proxy/sample"
	"testing"
)

func Test_Proxy(t *testing.T) {
	p := sample.NewPrinterProxy("Alice")
	t.Log("现在的名字是", p.GetPrintName(), "。")
	p.SetPrintName("Bob")
	t.Log("现在的名字是", p.GetPrintName(), "。")
	p.Print("Hello World")
}
