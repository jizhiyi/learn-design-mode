package a2

import (
	"testing"
)

func TestFileProperties_GetValue(t *testing.T) {
	fp := NewFileProperties()
	err := fp.ReadFromFile("file.txt")
	if err != nil {
		t.Fatal(err)
	}
	fp.SetValue("year", "2022")
	fp.SetValue("month", "10")
	fp.SetValue("day", "4")
	err = fp.WriteToFile("newfile.txt")
	if err != nil {
		t.Fatal(err)
	}
}
