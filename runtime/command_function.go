package runtime

import (
	"fmt"
	"strings"
)

type CommandFunctionDefine struct {
	Function *Function
}

func (c *CommandFunctionDefine) String() string {
	return fmt.Sprintf("CommandFunctionDefine[Function: \n  %v\n]", strings.Join(strings.Split(c.Function.String(), "\n"), "\n  "))
}

func (c *CommandFunctionDefine) Eval(stack *Stack, context *State) {
	context.DefineFunc(c.Function)
}

type CommandFunctionAccess struct {
	FunctionName string
}

func (c *CommandFunctionAccess) String() string {
	return fmt.Sprintf("CommandFunctionAccess[name: %s]", c.FunctionName)
}

func (c *CommandFunctionAccess) Eval(stack *Stack, context *State) {
	context.Err = UnImplemented.New("could not evaluate function access command")
}
