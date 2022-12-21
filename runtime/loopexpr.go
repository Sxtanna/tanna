package runtime

import (
	"fmt"
	"strings"

	"tanna/commons/coerce"
)

var LoopExpressionDeferred = &LoopExpressionDef{}

type LoopExpression interface {
	fmt.Stringer

	Init(stack *Stack, context *State)

	Eval(stack *Stack, context *State) bool
}

type LoopExpressionDef struct {
}

func (l *LoopExpressionDef) String() string {
	return "LoopExpressionDef"
}

func (l *LoopExpressionDef) Init(stack *Stack, context *State) {
	// nothing
}

func (l *LoopExpressionDef) Eval(stack *Stack, context *State) bool {
	return false
}

type LoopExpressionBit struct {
	Expr *Route
}

func (l *LoopExpressionBit) String() string {
	return fmt.Sprintf("LoopExpressionBit[\n  expr: \n    %s\n]", strings.Join(strings.Split(l.Expr.String(), "\n"), "\n    "))
}

func (l *LoopExpressionBit) Init(stack *Stack, context *State) {
	// nothing to initialize
}

func (l *LoopExpressionBit) Eval(stack *Stack, context *State) bool {
	if err := ExecuteRoute(stack, context, l.Expr); err != nil {
		return false
	}

	if v, ok := stack.Pull(); !ok {
		context.Err = StackIsEmpty.
			New("could not evaluate loop expression, stack was empty")
		return false
	} else if value, ok := v.(*Value); !ok {
		context.Err = InvalidValue.
			New("could not evaluate loop expression, pulled value isn't a tanna value")
		return false
	} else {
		return coerce.ToBoolean(value)
	}
}

type LoopExpressionRange struct {
	I *IntRangeIter
	D *DecRangeIter
	A *ArrRangeIter
}

func (l *LoopExpressionRange) String() string {
	return fmt.Sprintf("LoopExpressionRange[i: %v, d: %v, a: %v]", l.I, l.D, l.A)
}

func (l *LoopExpressionRange) Init(stack *Stack, context *State) {
	if v, ok := stack.Pull(); !ok {
		context.Err = StackIsEmpty.
			New("could not create number range, stack was empty")
		return
	} else if r, ok := v.(*Value); !ok {
		context.Err = InvalidValue.
			New("could not create number range, pulled value isn't a tanna value")
		return
	} else {
		switch r.Model {
		case RIInt:
			l.I = r.Value.(*IntRangeIter)
		case RIDec:
			l.D = r.Value.(*DecRangeIter)
		}
	}
}

func (l *LoopExpressionRange) Eval(stack *Stack, context *State) bool {
	if l.I != nil {
		if !l.I.Cont() {
			return false
		}

		stack.Push(&Value{
			Model: Int,
			Value: l.I.Next(),
		})
		return true
	}

	if l.D != nil {
		if !l.D.Cont() {
			return false
		}

		stack.Push(&Value{
			Model: Dec,
			Value: l.D.Next(),
		})
		return true
	}

	if l.A != nil {
		if !l.A.Cont() {
			return false
		}

		stack.Push(l.A.Next())
		return true
	}

	return false
}
