package runtime

import "github.com/joomcode/errorx"

var (
	RuntimeErrors = errorx.NewNamespace("runtime")

	UnImplemented = RuntimeErrors.NewType("unimplemented")

	StackIsEmpty = RuntimeErrors.NewType("stack_is_empty")

	InvalidValue = RuntimeErrors.NewType("invalid_value")

	PropertyNotFound = RuntimeErrors.NewType("property_not_found")

	FunctionNotFound = RuntimeErrors.NewType("function_not_found")

	LoopStates = RuntimeErrors.NewSubNamespace("loop")

	LoopStop = LoopStates.NewType("loop_stop")

	LoopCont = LoopStates.NewType("loop_cont")
)
