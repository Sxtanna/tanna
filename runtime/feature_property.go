package runtime

import (
	"fmt"
)

type Property struct {
	Name string

	Final bool
	Model Model
	Value *Value
}

func (p *Property) String() string {
	var word string
	if p.Final {
		word = "constant"
	} else {
		word = "variable"
	}

	var typeName string
	if p.Model == nil {
		typeName = ""
	} else {
		typeName = fmt.Sprintf(": %s", p.Model.Name())
	}

	return fmt.Sprintf("Property[%s %s%s]", word, p.Name, typeName)
}

func NewProperty(name string, final bool) *Property {
	return &Property{
		Name:  name,
		Final: final,
		Model: nil,
		Value: nil,
	}
}
