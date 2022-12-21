package runtime

type Scope struct {
	Name string

	Props map[string]*Property
	Funcs map[string]*Function

	Types map[string]Model
}

func NewScope(name string) *Scope {
	return &Scope{
		Name: name,

		Props: make(map[string]*Property),
		Funcs: make(map[string]*Function),

		Types: make(map[string]Model),
	}
}
