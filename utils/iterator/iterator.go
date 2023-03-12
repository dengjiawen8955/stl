package iterator

// ConstIterator 是一个默认的只读迭代器接口
type ConstIterator[T any] interface {
	// 是否有效
	IsValid() bool
	// 下一个迭代器
	Next() ConstIterator[T]
	// 当前元素的值
	Value() T
	// 克隆一个迭代器
	Clone() ConstIterator[T]
	// 比较两个迭代器是否相等
	Equal(other ConstIterator[T]) bool
}

// Iterator is an interface of mutable iterator
type Iterator[T any] interface {
	ConstIterator[T]
	SetValue(value T)
}

// ConstKvIterator is an interface of const key-value type iterator
type ConstKvIterator[K, V any] interface {
	ConstIterator[V]
	Key() K
}

// KvIterator is an interface of mutable key-value type iterator
type KvIterator[K, V any] interface {
	ConstKvIterator[K, V]
	SetValue(value V)
}

// ConstBidIterator is an interface of const bidirectional iterator
type ConstBidIterator[T any] interface {
	ConstIterator[T]
	Prev() ConstBidIterator[T]
}

// BidIterator is an interface of mutable bidirectional iterator
type BidIterator[T any] interface {
	ConstBidIterator[T]
	SetValue(value T)
}

// ConstKvBidIterator is an interface of const key-value type bidirectional iterator
type ConstKvBidIterator[K, V any] interface {
	ConstKvIterator[K, V]
	BidIterator[V]
}

// KvBidIterator is an interface of mutable key-value type bidirectional iterator
type KvBidIterator[K, V any] interface {
	KvIterator[K, V]
	BidIterator[V]
}

// RandomAccessIterator 支持随机访问的迭代器
type RandomAccessIterator[T any] interface {
	BidIterator[T]
	// 从当前位置移动offset个位置，正数向后移动，负数向前移动
	IteratorAt(position int) RandomAccessIterator[T]
	// 返回当前位置
	Position() int
}
