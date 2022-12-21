package words

import sets "github.com/deckarep/golang-set/v2"

// Word represents tanna's keywords
type Word string

const (
	Constant Word = "constant" // constant property
	Variable Word = "variable" // variable property
	Function Word = "function" // function
	Platform Word = "platform" // platform type resolution

	In Word = "in"
	By Word = "by"

	When Word = "when" // if statement
	Else Word = "else" // else/else-if statement
	Case Word = "case" // switch statement

	Loop Word = "loop" // looping statement
	Stop Word = "stop" // loop break statement
	Cont Word = "cont" // loop continue statement

	Range Word = "range"

	Cast Word = "cast" // type cast statement

	Sout Word = "sout" // print to console
	Read Word = "read" // read from console

	Push Word = "push" // push a value onto the current stack
	Pull Word = "pull" // pull a value from the current stack

	Class Word = "class" // object class
	Trait Word = "trait" // object trait
)

var Words = sets.NewThreadUnsafeSet[Word](Constant, Variable, Function, Platform, In, By, When, Else, Case, Loop, Stop, Cont, Range, Cast, Sout, Read, Push, Pull, Class, Trait)
var WordsAsStrings = sets.NewThreadUnsafeSet[string]("constant", "variable", "function", "platform", "in", "by", "when", "else", "case", "loop", "stop", "cont", "range", "cast", "sout", "read", "push", "pull", "class", "trait")
