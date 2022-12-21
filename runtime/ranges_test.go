package runtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IntRange(t *testing.T) {
	min := int64(0.0)
	max := int64(10.0)

	r := NewIntRange(min, max)

	assert.Equal(t,
		min,
		r.Min(), "range min should be %d", min)

	assert.Equal(t,
		max,
		r.Max(), "range max should be %d", max)
}

func Test_IntRangeIter(t *testing.T) {
	r := NewIntRange(0, 10)
	i := NewIntRangeIter(r)

	ex := int64(0)

	for i.Cont() {
		assert.Equal(t,
			ex,
			i.Next(), "range iter next should be %d", ex)

		ex += 1
	}
}

func Test_IntRangeIterWithStep(t *testing.T) {
	r := NewIntRange(0, 10)
	i := NewIntRangeIterWithStep(r, 2)

	ex := int64(0)

	for i.Cont() {
		assert.Equal(t,
			ex,
			i.Next(), "range iter next should be %d", ex)

		ex += 2
	}
}

func Test_DecRange(t *testing.T) {
	min := float64(0.0)
	max := float64(10.0)

	r := NewDecRange(min, max)

	assert.Equal(t,
		min,
		r.Min(), "range min should be %d", min)

	assert.Equal(t,
		max,
		r.Max(), "range max should be %d", max)

}

func Test_DecRangeIter(t *testing.T) {
	r := NewDecRange(0.0, 10.0)
	i := NewDecRangeIter(r)

	ex := float64(0)

	for i.Cont() {
		assert.Equal(t,
			ex,
			i.Next(), "range iter next should be %d", ex)

		ex += 1.0
	}
}

func Test_DecRangeIterWithStep(t *testing.T) {
	r := NewDecRange(0.0, 10.0)
	i := NewDecRangeIterWithStep(r, 2.0)

	ex := float64(0)

	for i.Cont() {
		assert.Equal(t,
			ex,
			i.Next(), "range iter next should be %d", ex)

		ex += 2.0
	}
}

func Test_ArrRange(t *testing.T) {
	arr := []*Value{
		{
			Model: Txt,
			Value: "H",
		},
		{
			Model: Txt,
			Value: "e",
		},
		{
			Model: Txt,
			Value: "l",
		},
		{
			Model: Txt,
			Value: "l",
		},
		{
			Model: Txt,
			Value: "o",
		},
	}

	r := NewArrRange(arr)

	assert.Equal(t,
		arr[0],
		r.Min(), "range min should be %s", "H")

	assert.Equal(t,
		arr[len(arr)-1],
		r.Max(), "range max should be %s", "o")

}

func Test_ArrRangeIter(t *testing.T) {
	arr := []*Value{
		{
			Model: Txt,
			Value: "H",
		},
		{
			Model: Txt,
			Value: "e",
		},
		{
			Model: Txt,
			Value: "l",
		},
		{
			Model: Txt,
			Value: "l",
		},
		{
			Model: Txt,
			Value: "o",
		},
	}

	r := NewArrRange(arr)
	i := NewArrRangeIter(r)

	index := 0

	for i.Cont() {
		assert.Equal(t,
			arr[index],
			i.Next(), "range iter next should be %s", arr[index])

		index = index + 1
	}
}
