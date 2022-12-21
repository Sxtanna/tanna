package runtime

import (
	"fmt"

	"tanna/commons/coerce"
)

var OperatorNot = new_op_bool("not", RTL, PR1, func(stack *Stack) bool {
	value, ok := stack.Pull()
	if !ok {
		panic(fmt.Errorf("could not pull value"))
	}

	bit := coerce.ToBoolean(value)

	return !bit
})

type operatorBool struct {
	Operator

	Function func(stack *Stack) bool
}

func new_op_bool(name string, associativity Associativity, precedence Precedence, function func(stack *Stack) bool) *Operator {
	operator := &operatorBool{Function: function}

	operator.Operator = Operator{
		Name:          name,
		Associativity: associativity,
		Precedence:    precedence,
		Function: func(stack *Stack) {
			stack.Push(&Value{
				Model: Bit,
				Value: operator.Function(stack),
			})
		},
	}

	return &operator.Operator
}
