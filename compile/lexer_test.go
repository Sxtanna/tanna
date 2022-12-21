package compile

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"tanna/commons/token"
)

const (
	lexer_funcLines = `function add(i0, i1) {



}`

	lexer_funcNoLines = `function add(i0, i1) {
}`

	lexer_code = `
	constant value0 = 10
	constant value1 = 20

	function add(a: Int, b: Int): Int {
		=> a + b
	}

	add(value0, value1)
`
)

type TokenDataComparison struct {
	Type token.Type
	Data string
}

func TestLexerPass0(t *testing.T) {
	lexer := &Lexer{
		code: lexer_funcLines,
	}

	lexer.pass0()

	assert.Equal(t, []rune(lexer_funcLines), lexer.dat0)
}

func TestLexerPass1(t *testing.T) {
	lexer := &Lexer{
		code: lexer_funcLines,
	}

	lexer.pass0()
	lexer.pass1()

	assert.Equal(t, []rune(lexer_funcNoLines), lexer.dat0)
}

func TestLexerPass2Numb(t *testing.T) {
	lexer := &Lexer{code: `0 1 12 1.5 .5 -1 -2.0 -.345`}

	lexer.RunLexer()
	lexer.pass4_remove_extra_values(true, false, false)

	assertTokens(t, lexer, []TokenDataComparison{
		{token.Numb, `0`},
		{token.Numb, `1`},
		{token.Numb, `12`},
		{token.Numb, `1.5`},
		{token.Numb, `0.5`},

		{token.Numb, `-1`},
		{token.Numb, `-2.0`},
		{token.Numb, `-0.345`},
	})
}

func TestLexerPass2Text(t *testing.T) {
	lexer := &Lexer{code: `'h'"hello""" "  "`}

	lexer.RunLexer()
	lexer.pass4_remove_extra_values(true, false, false)

	assertTokens(t, lexer, []TokenDataComparison{
		{token.Text, `'h'`},
		{token.Text, `"hello"`},
		{token.Text, `""`},
		{token.Text, `"  "`},
	})
}

func TestLexerPass2Word(t *testing.T) {
	lexer := &Lexer{code: `this class and that trait with one loop is true or false`}

	lexer.RunLexer()
	lexer.pass4_remove_extra_values(true, false, false)

	assertTokens(t, lexer, []TokenDataComparison{
		{token.Symb, `this`},
		{token.Word, `class`},
		{token.Symb, `and`},
		{token.Symb, `that`},
		{token.Word, `trait`},
		{token.Symb, `with`},
		{token.Symb, `one`},
		{token.Word, `loop`},
		{token.Symb, `is`},
		{token.Bool, `true`},
		{token.Symb, `or`},
		{token.Bool, `false`},
	})
}

func TestLexerPass2Oper(t *testing.T) {
	lexer := &Lexer{code: `+-*/<>!&|`}

	lexer.RunLexer()

	assertTokens(t, lexer, []TokenDataComparison{
		{token.Oper, `+`},
		{token.Oper, `-`},
		{token.Oper, `*`},
		{token.Oper, `/`},
		{token.Oper, `<`},
		{token.Oper, `>`},
		{token.Oper, `!`},
		{token.Oper, `&`},
		{token.Oper, `|`},
	})
}

func TestLexerPass2OperExt(t *testing.T) {
	lexer := &Lexer{code: `i++ a-- a *= 10`}

	lexer.RunLexer()
	lexer.pass4_remove_extra_values(true, false, false)

	assertTokens(t, lexer, []TokenDataComparison{
		{token.Symb, `i`},
		{token.Oper, `++`},

		{token.Symb, `a`},
		{token.Oper, `--`},

		{token.Symb, `a`},
		{token.Oper, `*=`},
		{token.Numb, `10`},
	})
}

func TestLexerPass2Comment(t *testing.T) {
	lexer := &Lexer{code: `class // comment here
// another comment
1 + 2`}

	lexer.RunLexer()
	lexer.pass4_remove_extra_values(true, false, false)

	assertTokens(t, lexer, []TokenDataComparison{
		{token.Word, `class`},
		{token.Numb, `1`},
		{token.Oper, `+`},
		{token.Numb, `2`},
	})
}

func TestLexerPass2Code(t *testing.T) {
	lexer := &Lexer{code: lexer_code}

	lexer.RunLexer()
	lexer.pass4_remove_extra_values(true, true, false)

	fmt.Printf("tokens length is %d\n", len(lexer.data))

	assertTokens(t, lexer, []TokenDataComparison{
		// properties
		{token.Word, `constant`},
		{token.Symb, `value0`},
		{token.Assign, `=`},
		{token.Numb, `10`},
		{token.Word, `constant`},
		{token.Symb, `value1`},
		{token.Assign, `=`},
		{token.Numb, `20`},

		// function
		{token.Word, `function`},
		{token.Symb, `add`},
		{token.ParenL, `(`},
		{token.Symb, `a`},
		{token.Typed, `:`},
		{token.Symb, `Int`},
		{token.Comma, `,`},
		{token.Symb, `b`},
		{token.Typed, `:`},
		{token.Symb, `Int`},
		{token.ParenR, `)`},
		{token.Typed, `:`},
		{token.Symb, `Int`},
		{token.BraceL, `{`},
		{token.Return, `=>`},
		{token.Symb, `a`},
		{token.Oper, `+`},
		{token.Symb, `b`},
		{token.BraceR, `}`},

		// function call
		{token.Symb, `add`},
		{token.ParenL, `(`},
		{token.Symb, `value0`},
		{token.Comma, `,`},
		{token.Symb, `value1`},
		{token.ParenR, `)`},
	})
}

func BenchmarkWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lexer := &Lexer{code: lexer_code, data: make([]*token.Data, 0, 50)}

		lexer.RunLexer()
	}
}

func assertToken(t *testing.T, lexer *Lexer, index int, tokenType token.Type, tokenData string) {
	_assertLexerDataHasIndex(t, lexer, index)

	_assertTokenTypeEqual(t, lexer, index, tokenType)
	_assertTokenDataEqual(t, lexer, index, tokenData)
}

func assertTokens(t *testing.T, lexer *Lexer, tokens []TokenDataComparison) {
	_assertLexerDataHasIndex(t, lexer, len(tokens)-1)

	for index, token := range tokens {
		_assertTokenTypeEqual(t, lexer, index, token.Type)
		_assertTokenDataEqual(t, lexer, index, token.Data)
	}
}

func assertTokenType(t *testing.T, lexer *Lexer, index int, tokenType token.Type) {
	_assertLexerDataHasIndex(t, lexer, index)

	_assertTokenTypeEqual(t, lexer, index, tokenType)
}

func assertTokenData(t *testing.T, lexer *Lexer, index int, tokenData string) {
	_assertLexerDataHasIndex(t, lexer, index)

	_assertTokenDataEqual(t, lexer, index, tokenData)
}

func _assertLexerDataHasIndex(t *testing.T, lexer *Lexer, index int) {
	assert.GreaterOrEqual(t, len(lexer.data), index+1, "lexer data does not contain index %d", index)
}

func _assertTokenTypeEqual(t *testing.T, lexer *Lexer, index int, tokenType token.Type) {
	assert.Equal(t, tokenType, lexer.data[index].Type, "token type at index %d should be `%v` but it is `%v`", index, tokenType, lexer.data[index].Type)
}

func _assertTokenDataEqual(t *testing.T, lexer *Lexer, index int, tokenData string) {
	assert.Equal(t, tokenData, lexer.data[index].Data, "token data at index %d should be `%v` but it is `%v`", index, tokenData, lexer.data[index].Data)
}
