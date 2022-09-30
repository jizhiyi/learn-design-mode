package iteratorversion2

// Aggregate 迭代器生成接口
type Aggregate[T any] interface {
	Iterator() Iterator[T]
}

// Iterator 迭代器接口
type Iterator[T any] interface {
	HasNext() bool
	Next() T
}
