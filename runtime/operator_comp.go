package runtime

import (
	"fmt"

	"tanna/commons/coerce"
)

var OperatorElse = new_op_comp("else", RTL, PR0, func(value0, value1 *Value) bool {
	b0 := coerce.ToBoolean(value0)
	b1 := coerce.ToBoolean(value1)

	return b0 || b1
})

var OperatorBoth = new_op_comp("both", RTL, PR0, func(value0, value1 *Value) bool {
	b0 := coerce.ToBoolean(value0)
	b1 := coerce.ToBoolean(value1)

	return b0 && b1
})

var OperatorSame = new_op_comp("same", RTL, PR0, func(value0, value1 *Value) bool {
	return value0.Equals(value1)
})

var OperatorDiff = new_op_comp("diff", RTL, PR0, func(value0, value1 *Value) bool {
	return !value0.Equals(value1)
})

type operatorComp struct {
	Operator

	Function func(value0, value1 *Value) bool
}

func new_op_comp(name string, associativity Associativity, precedence Precedence, function func(value0, value1 *Value) bool) *Operator {
	operator := &operatorComp{Function: function}

	operator.Operator = Operator{
		Name:          name,
		Associativity: associativity,
		Precedence:    precedence,
		Function: func(stack *Stack) {
			value0, ok := stack.Pull()
			if !ok {
				panic(fmt.Errorf("could not pull value 0"))
			}

			value1, ok := stack.Pull()
			if !ok {
				panic(fmt.Errorf("could not pull value 1"))
			}

			if v0, ok := value0.(*Value); !ok {
				panic(fmt.Errorf("could not coerce value 0 to tanna Value"))
			} else if v1, ok := value1.(*Value); !ok {
				panic(fmt.Errorf("could not coerce value 1 to tanna Value"))
			} else {
				stack.Push(&Value{
					Model: Bit,
					Value: operator.Function(v0, v1),
				})
			}
		},
	}

	return &operator.Operator
}
