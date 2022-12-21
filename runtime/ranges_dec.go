package runtime

import "tanna/commons/maths"

type DecRange struct {
	min, max float64
}

func NewDecRange(min, max float64) Range[float64] {
	return &DecRange{
		min: min,
		max: max,
	}
}

func (i *DecRange) Min() float64 {
	return i.min
}

func (i *DecRange) Max() float64 {
	return i.max
}

type DecRangeIter struct {
	Range[float64]

	step float64
	next float64
}

func NewDecRangeIter(target Range[float64]) RangeIter[float64] {
	return NewDecRangeIterWithStep(target, 1.0)
}

func NewDecRangeIterWithStep(target Range[float64], step float64) RangeIter[float64] {
	return &DecRangeIter{
		Range: target,

		step: step,
		next: target.Min(),
	}
}

func (i *DecRangeIter) Min() float64 {
	return i.Range.Min()
}

func (i *DecRangeIter) Max() float64 {
	return i.Range.Max()
}

func (i *DecRangeIter) Cont() bool {
	return i.next < i.Max()
}

func (i *DecRangeIter) Next() float64 {
	val := i.next

	i.next = maths.Min(i.Max(), i.next+i.step)

	return val
}
