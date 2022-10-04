package a2

type FileIO interface {
	ReadFromFile(filename string) error
	WriteToFile(filename string) error
	SetValue(key, value string)
	GetValue(key string) string
}
