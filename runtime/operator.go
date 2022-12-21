package runtime

import (
	sets "github.com/deckarep/golang-set/v2"
)

var Opers = sets.NewThreadUnsafeSet[rune]('+', '-', '*', '/', '<', '>', '!', '&', '|')

// Precedence Operators with higher precedence are evaluated first, regardless of associativity
type Precedence uint8

// Associativity Operators with the same precedence are grouped and evaluated based on associativity
type Associativity uint8

const (
	// NAN no associativity
	NAN Associativity = iota
	// LTR left-to-right associativity
	LTR
	// RTL right-to-left associativity
	RTL
)

const (
	PR0 Precedence = iota + 1
	PR1
	PR2
	PR3
	PR4
	PR5
	PR6
	PR7
	PR8
	PR9
)

type Operator struct {
	Name string

	Associativity
	Precedence

	Function func(stack *Stack)
}

var OperatorSOS = &Operator{
	Name:     "SOS",
	Function: func(stack *Stack) {},
}

var OperatorEOS = &Operator{
	Name:     "EOS",
	Function: func(stack *Stack) {},
}

func (o *Operator) String() string {
	switch o {
	case OperatorSOS:
		return "SoS"
	case OperatorEOS:
		return "EoS"
	case OperatorAdd:
		return "Add"
	case OperatorSub:
		return "Sub"
	case OperatorMul:
		return "Mul"
	case OperatorDiv:
		return "Div"

	case OperatorNot:
		return "Not"

	case OperatorElse:
		return "Else"
	case OperatorBoth:
		return "Both"
	case OperatorSame:
		return "Same"
	case OperatorDiff:
		return "Diff"
	}

	return "Unknown"
}
