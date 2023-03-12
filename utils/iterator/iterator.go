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

// Iterator 是一个可写的迭代器接口
type Iterator[T any] interface {
	ConstIterator[T]
	// 设置当前元素的值
	SetValue(value T)
}

// ConstKvIterator 是一个只读的键值对迭代器接口
type ConstKvIterator[K, V any] interface {
	// 迭代值
	ConstIterator[V]
	// 迭代键
	Key() K
}

// KvIterator 是一个可写的键值对迭代器接口
type KvIterator[K, V any] interface {
	// 迭代键, 值
	ConstKvIterator[K, V]
	// 设置值
	SetValue(value V)
}

// ConstBidIterator 是一个只读的双向迭代器接口, 可以向前或向后移动
type ConstBidIterator[T any] interface {
	// 上一个迭代器
	ConstIterator[T]
	// 下一个迭代器
	Prev() ConstBidIterator[T]
}

// BidIterator 是一个可写的双向迭代器接口, 可以向前或向后移动
type BidIterator[T any] interface {
	ConstBidIterator[T]
	// 设置当前元素的值
	SetValue(value T)
}

// ConstKvBidIterator 是一个只读的双向键值对迭代器接口, 可以向前或向后移动
type ConstKvBidIterator[K, V any] interface {
	// 迭代键值
	ConstKvIterator[K, V]
	BidIterator[V]
}

// KvBidIterator 是一个可写的双向键值对迭代器接口, 可以向前或向后移动
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
