package runtime

func Execute(cmds []Command) (stack *Stack, context *State, err error) {
	stack = NewStack()
	context = NewContext("global")

	if e := ExecuteRoute(stack, context, NewRoute(cmds)); e != nil {
		return stack, context, e
	}

	return
}

func ExecuteRoute(stack *Stack, context *State, route *Route) (err error) {
	here := route

	for here != nil {
		here.Eval(stack, context)

		// fmt.Printf("Executed %s, stack is now %s\n", reflect.TypeOf(here.Command), stack)

		if context.Err != nil {
			return context.Err
		}

		here = here.Next
	}

	return
}
