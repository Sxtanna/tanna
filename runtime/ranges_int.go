package runtime

import (
	"fmt"

	"tanna/commons/maths"
)

type IntRange struct {
	min, max int64
}

func (i *IntRange) String() string {
	return fmt.Sprintf("Range[%d:%d]", i.min, i.max)
}

func NewIntRange(min, max int64) Range[int64] {
	return &IntRange{
		min: min,
		max: max,
	}
}

func (i *IntRange) Min() int64 {
	return i.min
}

func (i *IntRange) Max() int64 {
	return i.max
}

type IntRangeIter struct {
	Range[int64]

	step int64
	next int64
}

func (i *IntRangeIter) String() string {
	return fmt.Sprintf("IntRangeIter[range: %s, step: %d]", i.Range, i.step)
}

func NewIntRangeIter(target Range[int64]) RangeIter[int64] {
	return NewIntRangeIterWithStep(target, 1)
}

func NewIntRangeIterWithStep(target Range[int64], step int64) RangeIter[int64] {
	return &IntRangeIter{
		Range: target,

		step: step,
		next: target.Min(),
	}
}

func (i *IntRangeIter) Min() int64 {
	return i.Range.Min()
}

func (i *IntRangeIter) Max() int64 {
	return i.Range.Max()
}

func (i *IntRangeIter) Cont() bool {
	return i.next < i.Max()
}

func (i *IntRangeIter) Next() int64 {
	val := i.next

	i.next = maths.Min(i.Max(), i.next+i.step)

	return val
}
