package token

// Type represents the type of the token
type Type uint8

const (
	_ Type = iota

	Numb // 0, 1.0, .5, -1, -2.3, -.32
	Text // 'a', "hello", `sup!`, "Hello $name", `Sup 1 + ${2 - value}`
	Bool // true, false

	Word // when, else, loop, case
	Oper // +, -, *, /
	Symb // {[a-zA-Z_][a-zA-Z0-9_]*} thisWorks, SoDoesThis | 0thisDoesNot

	Comma // ,
	Chain // .
	Space //  <- the space lol
	NLine // \n

	Point // &
	Deref // *

	Typed // methodName(): TypeHere
	Bound // TypeHere::[SuperTypeHere], platform::["NameOfPlatformType"]
	Range // 0..10

	Assign // =
	Return // =>

	BraceL // {
	BraceR // }
	BrackL // [
	BrackR // ]
	ParenL // (
	ParenR // )
)

func (t Type) String() string {
	switch t {
	case Numb:
		return "Numb"
	case Text:
		return "Text"
	case Bool:
		return "Bool"
	case Word:
		return "Word"
	case Oper:
		return "Oper"
	case Symb:
		return "Symb"
	case Comma:
		return "Comma"
	case Chain:
		return "Chain"
	case Space:
		return "Space"
	case NLine:
		return "NLine"
	case Point:
		return "Point"
	case Deref:
		return "Deref"
	case Typed:
		return "Typed"
	case Bound:
		return "Bound"
	case Range:
		return "Range"
	case Assign:
		return "Assign"
	case Return:
		return "Return"
	case BraceL:
		return "BraceL"
	case BraceR:
		return "BraceR"
	case BrackL:
		return "BrackL"
	case BrackR:
		return "BrackR"
	case ParenL:
		return "ParenL"
	case ParenR:
		return "ParenR"
	default:
		return "Unknown"
	}
}
