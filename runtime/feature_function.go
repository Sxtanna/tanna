package runtime

import (
	"fmt"
	"strings"
)

type Function struct {
	Name string

	Body *Route

	Accepts []*FunctionParameter
	Returns []*FunctionParameter
}

type FunctionParameter struct {
	Name  string
	Model Model
	None  *Value
}

func NewFunction(name string) *Function {
	return &Function{
		Name:    name,
		Body:    nil,
		Accepts: make([]*FunctionParameter, 0),
		Returns: make([]*FunctionParameter, 0),
	}
}

func NewFunctionParameter(name string, model Model) *FunctionParameter {
	return &FunctionParameter{
		Name:  name,
		Model: model,
		None:  nil,
	}
}

func (f *Function) AddAccepts(name string, model Model) *FunctionParameter {
	param := NewFunctionParameter(name, model)

	f.Accepts = append(f.Accepts, param)

	return param
}

func (f *Function) AddReturns(name string, model Model) *FunctionParameter {
	param := NewFunctionParameter(name, model)

	f.Returns = append(f.Returns, param)

	return param
}

func (f *Function) String() string {
	var accepts string

	if len(f.Accepts) == 0 {
		accepts = ""
	} else {
		builder := strings.Builder{}

		builder.WriteRune('(')

		for i, parameter := range f.Accepts {
			builder.WriteString(parameter.String())

			if i < len(f.Accepts)-1 {
				builder.WriteString(", ")
			}
		}

		builder.WriteRune(')')

		accepts = builder.String()
	}

	var returns string

	if len(f.Returns) == 0 {
		returns = ""
	} else {
		builder := strings.Builder{}

		builder.WriteString(": ")

		if len(f.Returns) == 1 {
			builder.WriteString(f.Returns[0].String())
		} else {
			builder.WriteRune('(')

			for i, parameter := range f.Returns {
				builder.WriteString(parameter.String())

				if i < len(f.Accepts)-1 {
					builder.WriteString(", ")
				}
			}

			builder.WriteRune(')')
		}

		returns = builder.String()
	}

	return fmt.Sprintf("Function[\n  fun %s%s%s {\n    %s\n  }\n]", f.Name, accepts, returns, strings.Join(strings.Split(f.Body.String(), "\n"), "\n    "))
}

func (f *FunctionParameter) String() string {
	var none string
	if f.None == nil {
		none = ""
	} else {
		none = fmt.Sprintf(" = %s", f.None)
	}

	return fmt.Sprintf("FP[%s: %s%s]", f.Name, f.Model.Name(), none)
}
