package runtime

type ArrRange struct {
	arr []*Value
}

func NewArrRange(arr []*Value) *ArrRange {
	return &ArrRange{
		arr: arr,
	}
}

func (a *ArrRange) Min() *Value {
	return a.arr[0]
}

func (a *ArrRange) Max() *Value {
	return a.arr[len(a.arr)-1]
}

type ArrRangeIter struct {
	*ArrRange

	index int
}

func NewArrRangeIter(target *ArrRange) *ArrRangeIter {
	return &ArrRangeIter{
		ArrRange: target,

		index: 0,
	}
}

func (i *ArrRangeIter) Min() *Value {
	return i.ArrRange.Min()
}

func (i *ArrRangeIter) Max() *Value {
	return i.ArrRange.Max()
}

func (i *ArrRangeIter) Cont() bool {
	return i.index < (len(i.ArrRange.arr) - 1)
}

func (i *ArrRangeIter) Next() *Value {
	val := i.ArrRange.arr[i.index]

	i.index = i.index + 1

	return val
}
