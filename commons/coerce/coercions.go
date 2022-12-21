package coerce

import (
	"fmt"
	"reflect"
)

func ToBoolean(val any) bool {
	switch t := (val).(type) {
	case bool:
		return t
	case int:
		// fallthrough
	case int8:
		// fallthrough
	case int16:
		// fallthrough
	case int32:
		// fallthrough
	case int64:
		if int64(t) == 0 {
			return false
		} else {
			return true
		}
	case float32:
		// fallthrough
	case float64:
		if float64(t) == 0.0 {
			return false
		} else {
			return true
		}
	case uint8:
		// fallthrough
	case uint16:
		// fallthrough
	case uint32:
		// fallthrough
	case uint64:
		if uint64(t) == 0 {
			return false
		} else {
			return true
		}
		/*case runtime.Value:
			return ToBoolean(t.Value)
		case *runtime.Value:
			return ToBoolean(t.Value)*/
	}

	panic(fmt.Errorf("could not coerce %v[%v] to bool", val, reflect.TypeOf(val)))
}

// ToFloat64 attempts to convert the generic value into a float64
//
// numbers use a direct coercion 'float64(val)'
// booleans are converted to '1.0' for true and '0.0' for false
// tanna Value types are extracted and recursively ran through this method
func ToFloat64(val any) float64 {
	switch t := (val).(type) {
	case int:
		return float64(t)
	case int8:
		return float64(t)
	case int16:
		return float64(t)
	case int32:
		return float64(t)
	case int64:
		return float64(t)
	case float32:
		return float64(t)
	case float64:
		return float64(t)
	case uint8:
		return float64(t)
	case uint16:
		return float64(t)
	case uint32:
		return float64(t)
	case uint64:
		return float64(t)
	case bool:
		if t {
			return 1.0
		} else {
			return 0.0
		}
		/*case runtime.Value:
			return ToFloat64(t.Value)
		case *runtime.Value:
			return ToFloat64(t.Value)*/
	}

	panic(fmt.Errorf("could not coerce %v[%v] to float64", val, reflect.TypeOf(val)))
}

// ToCommonNumberType attempts to convert the input value to one of tanna's common number types Int aka. 'int64' or Dec aka. 'float64'
func ToCommonNumberType(val float64, origin0, origin1 any) any {
	origin0Type := reflect.TypeOf(origin0)
	origin1Type := reflect.TypeOf(origin1)

	origin0TypeName := origin0Type.Name()
	origin1TypeName := origin1Type.Name()

	if origin0TypeName == "int" || origin1TypeName == "int" ||
		origin0TypeName == "int8" || origin1TypeName == "int8" ||
		origin0TypeName == "int16" || origin1TypeName == "int16" ||
		origin0TypeName == "int32" || origin1TypeName == "int32" ||
		origin0TypeName == "int64" || origin1TypeName == "int64" ||
		origin0TypeName == "uint" || origin1TypeName == "uint" ||
		origin0TypeName == "uint8" || origin1TypeName == "uint8" ||
		origin0TypeName == "uint16" || origin1TypeName == "uint16" ||
		origin0TypeName == "uint32" || origin1TypeName == "uint32" ||
		origin0TypeName == "uint64" || origin1TypeName == "uint64" {
		return int64(val)
	} else if origin0TypeName == "float32" ||
		origin1TypeName == "float32" ||
		origin0TypeName == "float64" ||
		origin1TypeName == "float64" {
		return val
	} else if origin0TypeName == "bool" ||
		origin1TypeName == "bool" {
		if val != 0.0 {
			return true
		} else {
			return false
		}
	}

	panic(fmt.Errorf("could not coerce %v => %v|%v", val, origin0Type, origin1Type))
}
