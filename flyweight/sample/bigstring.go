package sample

type BigString struct {
	bigChars []*BigChar
}

func (b *BigString) Print() {
	for _, char := range b.bigChars {
		char.print()
	}
}

func NewBigString(str string) *BigString {
	bigString := &BigString{}
	factory := GetFactoryInstance()
	for _, ch := range str {
		bigString.bigChars = append(bigString.bigChars, factory.GetBigChar(byte(ch)))
	}
	return bigString
}
