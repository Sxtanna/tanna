package token

import "fmt"

// Data represents a singularly lexed token
type Data struct {
	// the type of the token
	Type Type
	// the string data of the token
	Data string
	// which line the token was on
	Line int
	// which character on the line the token was at
	Char int
}

func (t *Data) String() string {
	return fmt.Sprintf("Token[%v:`%s` | %d:%d]", t.Type, t.Data, t.Line, t.Char)
}
