package a2

import (
	"fmt"
	"io"
	"strings"
)

type Properties struct {
	data map[string]string
}

func NewProperties() *Properties {
	return &Properties{data: make(map[string]string)}
}

func (p *Properties) load(r io.Reader) error {
	if r == nil {
		return fmt.Errorf("reader is nil")
	}
	for {
		var line string
		_, err := fmt.Fscanf(r, "%s\n", &line)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		strs := strings.Split(line, "=")
		if len(strs) != 2 {
			return fmt.Errorf("format not match")
		}
		p.data[strs[0]] = strs[1]
	}
	return nil
}

func (p *Properties) Store(w io.Writer) error {
	if w == nil {
		return fmt.Errorf("writer is nil")
	}
	for key, val := range p.data {
		_, err := fmt.Fprintf(w, "%s=%s\n", key, val)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Properties) SetProperty(key, value string) {
	p.data[key] = value
}

func (p *Properties) GetProperty(key string) string {
	return p.data[key]
}
