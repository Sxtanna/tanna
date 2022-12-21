package runtime

import (
	"fmt"
	"reflect"
)

var (
	ValueNil = &Value{
		Model: Nil,
		Value: nil,
	}
)

type Value struct {
	Model Model
	Value any
}

func (v *Value) String() string {
	return fmt.Sprintf("Value[%s]::[%v]", v.Model.Name(), v.Value)
}

func (v *Value) Equals(other any) bool {
	if o, ok := other.(*Value); !ok {
		return false
	} else if !v.Model.Same(o.Model) {
		return false
	} else if !reflect.DeepEqual(v.Value, o.Value) {
		return false
	}

	return true
}
