package a2

import "os"

type FileProperties struct {
	properties *Properties
}

func NewFileProperties() *FileProperties {
	return &FileProperties{properties: NewProperties()}
}

func (fp *FileProperties) ReadFromFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	return fp.properties.load(f)
}

func (fp *FileProperties) WriteToFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	return fp.properties.Store(f)
}

func (fp *FileProperties) SetValue(key string, value string) {
	fp.properties.SetProperty(key, value)
}

func (fp *FileProperties) GetValue(key string) string {
	return fp.properties.GetProperty(key)
}
