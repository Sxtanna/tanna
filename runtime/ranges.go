package runtime

type Range[T any] interface {
	Min() T
	Max() T
}

type RangeIter[T any] interface {
	Range[T]

	Cont() bool
	Next() T
}
