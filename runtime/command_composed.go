package runtime

import (
	"fmt"
)

type CommandOperator struct {
	Oper *Operator
}

func (c *CommandOperator) String() string {
	return fmt.Sprintf("CommandOperator[Oper: %v]", c.Oper)
}

func (c *CommandOperator) Eval(stack *Stack, context *State) {
	c.Oper.Function(stack)
}
