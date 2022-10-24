package sample

type BigCharFactory struct {
	charPool map[byte]*BigChar
}

var bigCharFactory *BigCharFactory

func GetFactoryInstance() *BigCharFactory {
	if bigCharFactory == nil {
		bigCharFactory = &BigCharFactory{}
		bigCharFactory.charPool = make(map[byte]*BigChar)
	}
	return bigCharFactory
}

func (b *BigCharFactory) GetBigChar(char byte) *BigChar {
	if char, ok := b.charPool[char]; ok {
		return char
	}
	return NewBigChar(char)
}
