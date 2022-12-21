package compile

import (
	"github.com/joomcode/errorx"
	"tanna/commons/token"
)

var (
	// ParserErrors parser error namespace
	ParserErrors = errorx.NewNamespace("parser")

	// TokenTypeOutOfPlace an error that occurs due to a token having a type in a place it shouldn't be
	TokenTypeOutOfPlace = ParserErrors.NewType("token_type_out_of_place")

	// TokenDataOutOfPlace an error that occurs due to a token having data in a place it shouldn't be
	TokenDataOutOfPlace = ParserErrors.NewType("token_data_out_of_place")

	// TokenRequiredMissing an error that occurs when an expected token is missing from where it should be
	TokenRequiredMissing = ParserErrors.NewType("token_required_missing")

	// ShuntingYardEmptyStack an error that occurs when executing the shunting yard algorithm results in an empty stack after flushing
	ShuntingYardEmptyStack = ParserErrors.NewType("empty_stack_after_flush")
)

func MakeTokenRequiredMissing(name string, expected token.Type, actual *token.Data, reason string) error {
	return TokenRequiredMissing.New("could not parse '%s': found '%v' expected to find `%v` token type: %s", name, actual, expected, reason)
}
