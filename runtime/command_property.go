package runtime

import (
	"fmt"
)

type CommandPropertyDefine struct {
	Property *Property
}

func (c *CommandPropertyDefine) String() string {
	return fmt.Sprintf("CommandPropertyDefine[Property: %v]", c.Property)
}

func (c *CommandPropertyDefine) Eval(stack *Stack, context *State) {
	context.DefineProp(c.Property)
}

type CommandPropertyAssign struct {
	PropertyName string
}

func (c *CommandPropertyAssign) String() string {
	return fmt.Sprintf("CommandPropertyAssign[PropertyName: %s]", c.PropertyName)
}

func (c *CommandPropertyAssign) Eval(stack *Stack, context *State) {
	property := context.LocateProp(c.PropertyName, -1)
	if property == nil {
		context.Err = PropertyNotFound.
			New("could not assign property %s, it was not found within the context", c.PropertyName)
		return
	}

	dat, ok := stack.Pull()
	if !ok {
		context.Err = StackIsEmpty.
			New("could not assign property %s, the stack is empty", c.PropertyName)
		return // stack was empty
	}

	value, ok := dat.(*Value)
	if !ok {
		context.Err = InvalidValue.
			New("could not assign property %s, the resolved data %s is not an assignable value", c.PropertyName, dat)
		return // not a valid assignable value
	}

	// todo: type assertions

	property.Value = value
}

type CommandPropertyAccess struct {
	PropertyName string
}

func (c *CommandPropertyAccess) String() string {
	return fmt.Sprintf("CommandPropertyAccess[PropertyName: %s]", c.PropertyName)
}

func (c *CommandPropertyAccess) Eval(stack *Stack, context *State) {
	property := context.LocateProp(c.PropertyName, -1)
	if property == nil {
		context.Err = PropertyNotFound.
			New("could not access property %s, it was not found within the context", c.PropertyName)
		return
	}

	value := property.Value
	if value == nil {
		stack.Push(ValueNil)
	} else {
		stack.Push(value)
	}
}

type CommandPropertyNilify struct {
	PropertyName string
}

func (c *CommandPropertyNilify) String() string {
	return fmt.Sprintf("CommandPropertyNilify[PropertyName: %s]", c.PropertyName)
}

func (c *CommandPropertyNilify) Eval(stack *Stack, context *State) {
	property := context.LocateProp(c.PropertyName, -1)
	if property == nil {
		context.Err = PropertyNotFound.
			New("could not nilify property %s, it was not found within the context", c.PropertyName)
		return
	}

	property.Value = ValueNil
}
