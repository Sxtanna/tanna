package runtime

import (
	"io"
	"os"
)

type State struct {
	Name string

	Scope *Stack

	Reader io.Reader
	Writer io.Writer

	Err error
}

func NewContext(name string) *State {
	context := &State{
		Name: name,

		Scope: NewStack(),

		Reader: os.Stdin,
		Writer: os.Stdout,
	}

	context.Enter(NewScope(name))

	// for _, builtIn := range BuiltIns {
	// 	context.DefineType(builtIn)
	// }

	return context
}

func (c *State) Enter(scope *Scope) {
	c.Scope.Push(scope)
}

func (c *State) Leave() *Scope {
	if scope, ok := c.Scope.Pull(); ok {
		return scope.(*Scope)
	}

	return nil
}

func (c *State) CurrentScope() *Scope {
	return (c.Scope.Peek()).(*Scope)
}

func (c *State) DefineProp(p *Property) {
	if scope := c.CurrentScope(); scope != nil {
		scope.Props[p.Name] = p
	}
}

func (c *State) DefineFunc(f *Function) {
	if scope := c.CurrentScope(); scope != nil {
		scope.Funcs[f.Name] = f
	}
}

func (c *State) DefineType(t Model) {
	if scope := c.CurrentScope(); scope != nil {
		scope.Types[t.Name()] = t
	}
}

func (c *State) LocateProp(name string, depth int) *Property {

	for i := len(*c.Scope) - 1; i >= 0; i-- {
		scope := ((*c.Scope)[i]).(*Scope)
		if scope == nil {
			continue
		}

		if prop, ok := scope.Props[name]; ok {
			return prop
		}

		if depth != -1 {
			depth--

			if depth <= 0 {
				break
			}
		}
	}

	return nil
}

func (c *State) LocateFunc(name string) *Function {

	for i := len(*c.Scope) - 1; i >= 0; i-- {
		scope := ((*c.Scope)[i]).(*Scope)
		if scope == nil {
			continue
		}

		if f, ok := scope.Funcs[name]; ok {
			return f
		}
	}

	return nil
}

func (c *State) LocateType(name string) Model {

	for i := len(*c.Scope) - 1; i >= 0; i-- {
		scope := ((*c.Scope)[i]).(*Scope)
		if scope == nil {
			continue
		}

		if t, ok := scope.Types[name]; ok {
			return t
		}
	}

	return nil
}
