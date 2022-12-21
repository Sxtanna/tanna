package runtime

import (
	"fmt"
	"log"
	"strings"

	"github.com/joomcode/errorx"
	"tanna/commons/coerce"
	"tanna/compile/utils"
)

type CommandBoundary struct {
}

func (c *CommandBoundary) String() string {
	return "CommandBoundary"
}

func (c *CommandBoundary) Eval(stack *Stack, context *State) {
	// do nothing
}

type CommandLiteral struct {
	Value *Value
}

func (c *CommandLiteral) String() string {
	return fmt.Sprintf("CommandLiteral[value: %v]", c.Value)
}

func (c *CommandLiteral) Eval(stack *Stack, context *State) {
	stack.Push(c.Value)
}

type CommandRoute struct {
	Route *Route
}

func (c *CommandRoute) String() string {
	return fmt.Sprintf("CommandRoute[route: \n  %s\n]", strings.Join(strings.Split(c.Route.String(), "\n"), "\n  "))
}

func (c *CommandRoute) Eval(stack *Stack, context *State) {
	_ = ExecuteRoute(stack, context, c.Route)
}

type CommandTuple struct {
	Size int
}

func (c *CommandTuple) String() string {
	return fmt.Sprintf("CommandTuple[size: %d]", c.Size)
}

func (c *CommandTuple) Eval(stack *Stack, context *State) {
	values := make([]*Value, 0)

	for i := 0; i < c.Size; i++ {
		if dat, ok := stack.Pull(); ok {
			if value, ok := dat.(*Value); !ok {
				context.Err = InvalidValue.
					New("could not resolve tuple index %d, the resolved data %s is not an assignable value", i, dat)
				return
			} else {
				values = append(values, value)
			}
		}
	}

	models := make([]Model, 0)
	for _, value := range values {
		models = append(models, value.Model)
	}

	resolvedModel := NewModelTuple(models)

	tuple := context.LocateType(resolvedModel.Name())
	if tuple == nil {
		tuple = resolvedModel
		context.DefineType(tuple)
	}

	utils.ReverseSlice(values)

	stack.Push(&Value{
		Model: tuple,
		Value: values,
	})
}

type CommandSout struct {
	AppendNewLine bool
	ReturnToStack bool
}

func (c *CommandSout) String() string {
	return fmt.Sprintf("CommandSout[appendNewLine: %t, returnToStack: %t]", c.AppendNewLine, c.ReturnToStack)
}

func (c *CommandSout) Eval(stack *Stack, context *State) {
	if value, ok := stack.Pull(); !ok {
		context.Err = StackIsEmpty.
			New("could not print value to output, stack was empty")
		return
	} else {
		if !c.AppendNewLine {
			log.Print(value)
		} else {
			log.Println(value)
		}

		if c.ReturnToStack {
			stack.Push(value)
		}
	}
}

type CommandPull struct {
	PullValueFromStack bool
}

func (c *CommandPull) String() string {
	return "CommandPull"
}

func (c *CommandPull) Eval(stack *Stack, context *State) {
	if value := stack.Peek(); value == nil {
		context.Err = StackIsEmpty.
			New("could not execute stack pull, stack is empty")
		return
	}

	if c.PullValueFromStack {
		stack.Pull()
	}
}

type CommandLoop struct {
	Vars *Route
	Body *Route
	Expr LoopExpression
}

func (c *CommandLoop) String() string {
	var expr string

	if c.Expr == nil {
		expr = "None"
	} else {
		expr = fmt.Sprintf("\n    %s", strings.Join(strings.Split(c.Expr.String(), "\n"), "\n    "))
	}

	return fmt.Sprintf("CommandLoop[\n  expr: %s\n  body: \n    %s\n]",
		expr,
		strings.Join(strings.Split(c.Body.String(), "\n"), "\n    "))
}

func (c *CommandLoop) Eval(stack *Stack, context *State) {

	if c.Vars != nil && c.Vars != None {
		context.Enter(NewScope("loop_vars"))

		if err := ExecuteRoute(stack, context, c.Vars); err != nil {
			context.Leave()
			return
		}
	}

	if c.Expr != nil {
		c.Expr.Init(stack, context)
	}

	for {
		pass := true

		if c.Expr != nil {
			pass = c.Expr.Eval(stack, context)
		}

		if !pass {
			break
		}

		if c.Body != nil && c.Body != None {
			context.Enter(NewScope("loop_body"))

			if err := ExecuteRoute(stack, context, c.Body); err == nil {
				continue
			} else {
				if errorx.IsOfType(err, LoopCont) {
					context.Err = nil
					continue
				}

				context.Leave()

				if errorx.IsOfType(err, LoopStop) {
					context.Err = nil
				}

				break
			}
		}
	}

	if c.Vars != nil && c.Vars != None {
		context.Leave()
	}
}

type CommandLoopStop struct {
	FullStop bool
}

func (c *CommandLoopStop) String() string {
	return fmt.Sprintf("CommandLoopStop[full_stop: %t]", c.FullStop)
}

func (c *CommandLoopStop) Eval(stack *Stack, context *State) {
	if c.FullStop {
		context.Err = LoopStop.NewWithNoMessage()
	} else {
		context.Err = LoopCont.NewWithNoMessage()
	}
}

type CommandWhen struct {
	Expr *Route
	Pass *Route
	Fail *Route
}

func (c *CommandWhen) String() string {
	return fmt.Sprintf("CommandWhen[\nexpr: \n  %s\npass: \n  %s\nfail: \n  %s\n]",
		strings.Join(strings.Split(c.Expr.String(), "\n"), "\n  "),
		strings.Join(strings.Split(c.Pass.String(), "\n"), "\n  "),
		strings.Join(strings.Split(c.Fail.String(), "\n"), "\n  "))
}

func (c *CommandWhen) Eval(stack *Stack, context *State) {
	err := ExecuteRoute(stack, context, c.Expr)
	if err != nil {
		return
	}

	var state bool

	if v, ok := stack.Pull(); !ok {
		context.Err = StackIsEmpty.
			New("could not evaluate when expression, stack was empty")
		return
	} else if value, ok := v.(*Value); !ok {
		context.Err = InvalidValue.
			New("could not evaluate when expression, pulled value isn't a tanna value")
		return
	} else {
		state = coerce.ToBoolean(value)
	}

	if state {
		_ = ExecuteRoute(stack, context, c.Pass)
	} else {
		_ = ExecuteRoute(stack, context, c.Fail)
	}
}

type CommandStackPush struct {
	Value any
}

func (c *CommandStackPush) String() string {
	return fmt.Sprintf("CommandStackPush[value: %s]", c.Value)
}

func (c *CommandStackPush) Eval(stack *Stack, context *State) {
	stack.Push(c.Value)
}

type CommandMakeRange struct {
	Num bool
	Arr bool
}

func (c *CommandMakeRange) String() string {
	return fmt.Sprintf("CommandMakeRange[num: %t, arr: %t]", c.Num, c.Arr)
}

func (c *CommandMakeRange) Eval(stack *Stack, context *State) {
	if c.Num {
		var min *Value
		var max *Value

		if v, ok := stack.Pull(); !ok {
			context.Err = StackIsEmpty.
				New("could not create number range, stack was empty")
			return
		} else if r, ok := v.(*Value); !ok {
			context.Err = InvalidValue.
				New("could not create number range, pulled value isn't a tanna value")
			return
		} else {
			max = r
		}

		if v, ok := stack.Pull(); !ok {
			context.Err = StackIsEmpty.
				New("could not create number range, stack was empty")
			return
		} else if r, ok := v.(*Value); !ok {
			context.Err = InvalidValue.
				New("could not create number range, pulled value isn't a tanna value")
			return
		} else {
			min = r
		}

		if !min.Model.Same(max.Model) {
			context.Err = InvalidValue.
				New("could not create number range, pulled values aren't the same type")
		} else {
			minNumber := min.Value
			maxNumber := max.Value

			switch min.Model {
			case Int:
				intMin := minNumber.(int64)
				intMax := maxNumber.(int64)

				stack.Push(&Value{
					Model: RInt,
					Value: NewIntRange(intMin, intMax),
				})
			case Dec:
				decMin := minNumber.(float64)
				decMax := maxNumber.(float64)

				stack.Push(&Value{
					Model: RDec,
					Value: NewDecRange(decMin, decMax),
				})
			default:
				context.Err = InvalidValue.
					New("could not create number range, pulled values aren't rangeable")
			}
		}
	}

}

type CommandMakeRangeIter struct {
}

func (c *CommandMakeRangeIter) String() string {
	return "CommandMakeRangeIter"
}

func (c *CommandMakeRangeIter) Eval(stack *Stack, context *State) {
	var step *Value

	if v, ok := stack.Pull(); !ok {
		context.Err = StackIsEmpty.
			New("could not create number range, stack was empty")
		return
	} else if r, ok := v.(*Value); !ok {
		context.Err = InvalidValue.
			New("could not create number range, pulled value isn't a tanna value")
		return
	} else {
		step = r
	}

	if v, ok := stack.Pull(); !ok {
		context.Err = StackIsEmpty.
			New("could not create range iter, stack was empty")
		return
	} else if r, ok := v.(*Value); !ok {
		context.Err = InvalidValue.
			New("could not create range iter, pulled value isn't a tanna value")
		return
	} else {
		switch r.Model {
		case RInt:
			stack.Push(&Value{
				Model: RIInt,
				Value: NewIntRangeIterWithStep(r.Value.(Range[int64]), step.Value.(int64)),
			})
		case RDec:
			stack.Push(&Value{
				Model: RIDec,
				Value: NewDecRangeIterWithStep(r.Value.(Range[float64]), step.Value.(float64)),
			})
		}
	}
}
