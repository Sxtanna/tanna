package runtime

import (
	"strings"
)

const split_style = false

var None = &Route{Command: CommandNone}

type Route struct {
	Command Command

	Prev *Route
	Next *Route
}

func NewRoute(cmds []Command) *Route {
	if len(cmds) == 0 {
		return None
	}

	routes := make([]*Route, 0)

	for _, cmd := range cmds {
		routes = append(routes, &Route{Command: cmd})
	}

	for i := range routes {
		route := routes[i]

		if i > 0 {
			route.Prev = routes[i-1]
		}
		if i < len(routes)-1 {
			route.Next = routes[i+1]
		}
	}

	return routes[0]
}

func (r *Route) Eval(stack *Stack, context *State) {
	if r.Command != nil {
		r.Command.Eval(stack, context)
	}
}

func (r *Route) String() string {
	if r == nil {
		return None.String()
	}

	unwrapped := r.unwrap()

	if split_style {
		unwrapped = unwrapped[1:]
	}

	builder := strings.Builder{}

	builder.WriteString("Route[")
	builder.WriteRune('\n')

	if split_style {
		builder.WriteString("  command: ")
		builder.WriteRune('\n')
		builder.WriteString("    ")
		builder.WriteString(strings.Join(strings.Split(r.Command.String(), "\n"), "\n    "))
		builder.WriteRune('\n')
		builder.WriteString("  wrapped: ")
		builder.WriteRune('\n')
	} else {
		builder.WriteString("  route: ")
		builder.WriteRune('\n')
	}

	for i, command := range unwrapped {
		builder.WriteString("    ")
		builder.WriteString(strings.Join(strings.Split(command.String(), "\n"), "\n    "))

		if i < len(unwrapped)-1 {
			builder.WriteRune('\n')
		}
	}

	builder.WriteRune('\n')
	builder.WriteString("]")

	return builder.String()
}

func (r *Route) unwrap() []Command {
	cmds := make([]Command, 0)

	here := r
	for here != nil {
		cmds = append(cmds, here.Command)
		here = here.Next
	}

	return cmds
}
