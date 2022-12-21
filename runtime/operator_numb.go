package runtime

import (
	"fmt"
	"reflect"

	"tanna/commons/coerce"
)

var OperatorAdd = new_op_numb("add", LTR, PR3, func(numb0, numb1 float64) float64 { return numb0 + numb1 })

var OperatorSub = new_op_numb("sub", LTR, PR3, func(numb0, numb1 float64) float64 { return numb0 - numb1 })

var OperatorMul = new_op_numb("mul", LTR, PR4, func(numb0, numb1 float64) float64 { return numb0 * numb1 })

var OperatorDiv = new_op_numb("div", LTR, PR4, func(numb0, numb1 float64) float64 { return numb0 / numb1 })

type operatorNumb struct {
	Operator

	Function func(numb0, numb1 float64) float64
}

func new_op_numb(name string, associativity Associativity, precedence Precedence, function func(numb0, numb1 float64) float64) *Operator {

	operator := &operatorNumb{Function: function}

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

			value := coerce.ToCommonNumberType(operator.Function(
				coerce.ToFloat64(value1), coerce.ToFloat64(value0)),
				ExtractActualValue(value0), ExtractActualValue(value1))

			model := tannaNumberType(value)

			stack.Push(&Value{
				Model: model,
				Value: value,
			})
		},
	}

	return &operator.Operator
}

func tannaNumberType(val any) Model {
	var name string

	switch val.(type) {
	case int:
		name = "Int"
	case int8:
		name = "Int"
	case int16:
		name = "Int"
	case int32:
		name = "Int"
	case int64:
		name = "Int"
	case float32:
		name = "Dec"
	case float64:
		name = "Dec"
	case uint8:
		name = "Int"
	case uint16:
		name = "Int"
	case uint32:
		name = "Int"
	case uint64:
		name = "Int"
	default:
		panic(fmt.Errorf("could not find type of %v | %v", val, reflect.TypeOf(val)))
	}

	return FindBuiltInModel(name)
}

func ExtractActualValue(value any) any {
	for {
		v, ok := value.(*Value)
		if !ok {
			return value
		}

		value = v.Value
	}
}
