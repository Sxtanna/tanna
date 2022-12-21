package runtime

import (
	"fmt"
)

var CommandNone Command = &commandNone{}

type Command interface {
	fmt.Stringer
	Eval(stack *Stack, context *State)
}

type commandNone struct {
}

func (c *commandNone) String() string {
	return "None"
}

func (c *commandNone) Eval(stack *Stack, context *State) {

}
