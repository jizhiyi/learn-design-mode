package iteratorversion1

// Aggregate 迭代器生成接口
type Aggregate interface {
	Iterator() Iterator
}

// Iterator 迭代器接口
type Iterator interface {
	HasNext() bool
	Next() interface{}
}
