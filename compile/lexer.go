package compile

import (
	"fmt"
	"strings"
	"unicode"

	"tanna/commons/token"
	"tanna/commons/words"
	"tanna/runtime"
)

var lexer_augmenters = [][]*token.Augmenter{
	{
		{
			BackTrack: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Numb,
					Data: fmt.Sprintf("0.%s", here.Data),
				}}
			},
			PrevMatch: func(data *token.Data) bool {
				return data.Type == token.Chain
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Numb && !strings.ContainsRune(data.Data, '.')
			},
		},
		{
			SkipCount: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Oper,
					Data: "++",
				}}
			},
			PrevMatch: func(data *token.Data) bool {
				return data.Type == token.Symb
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "+"
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "+"
			},
		},
		{
			SkipCount: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Oper,
					Data: "--",
				}}
			},
			PrevMatch: func(data *token.Data) bool {
				return data.Type == token.Symb
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "-"
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "-"
			},
		},
		{
			SkipCount: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Oper,
					Data: "&&",
				}}
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "&"
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "&"
			},
		},
		{
			SkipCount: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Oper,
					Data: "||",
				}}
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "|"
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "|"
			},
		},
		{
			SkipCount: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Bound,
					Data: "::",
				}}
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Typed
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Typed
			},
		},
		{
			SkipCount: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Return,
					Data: "=>",
				}}
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Assign
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == ">"
			},
		},
		{
			SkipCount: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Oper,
					Data: "+=",
				}}
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "+"
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Assign
			},
		},
		{
			SkipCount: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Oper,
					Data: "-=",
				}}
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "-"
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Assign
			},
		},
		{
			SkipCount: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Oper,
					Data: "*=",
				}}
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "*"
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Assign
			},
		},
		{
			SkipCount: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Oper,
					Data: "/=",
				}}
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "/"
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Assign
			},
		},
		{
			SkipCount: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Oper,
					Data: "==",
				}}
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Assign
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Assign
			},
		},
		{
			SkipCount: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Oper,
					Data: "!=",
				}}
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "!"
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Assign
			},
		},
		{
			SkipCount: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Oper,
					Data: ">=",
				}}
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == ">"
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Assign
			},
		},
		{
			SkipCount: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Oper,
					Data: "<=",
				}}
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "<"
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Assign
			},
		}},
	{
		{
			BackTrack: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Numb,
					Data: fmt.Sprintf("-%s", here.Data),
				}}
			},
			PrevMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "-"
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Numb
			},
		},
		{
			BackTrack: 1,
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{{
					Type: token.Oper,
					Data: "++",
				}}
			},
			PrevMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "+"
			},
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "+"
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Symb
			},
		},
	},
}

type Lexer struct {
	code string
	dat0 []rune
	data []*token.Data
	idx2 int

	line int
	char int
}

func NewLexer(code string) *Lexer {
	return &Lexer{code: code}
}

func (l *Lexer) addToken(tokenType token.Type, tokenData string) {
	l.data = append(l.data, &token.Data{
		Type: tokenType,
		Data: tokenData,
		Line: l.line,
		Char: l.char,
	})
}

func (l *Lexer) RunLexer() {
	l.pass0()
	l.pass1()
	l.pass2()
	l.pass3()
	l.pass4()
}

func (l *Lexer) LexerData() []*token.Data {
	return l.data
}

// convert string to runes
func (l *Lexer) pass0() {
	l.dat0 = []rune(l.code)
}

// remove extra new lines
func (l *Lexer) pass1() {
	for i := 0; i < len(l.dat0); i++ {
		if !unicode.IsSpace(l.dat0[i]) {
			break
		} else {
			l.dat0 = append(l.dat0[:i], l.dat0[i+1:]...)
			i--
		}
	}

	for i := len(l.dat0) - 1; i >= 0; i-- {
		if !unicode.IsSpace(l.dat0[i]) {
			break
		} else {
			l.dat0 = append(l.dat0[:i], l.dat0[i+1:]...)
		}
	}

	i := 1
	p := l.dat0[0]

	for i < len(l.dat0) {
		n := l.dat0[i]
		if p == '\n' && n == '\n' {
			l.dat0 = append(l.dat0[:i], l.dat0[i+1:]...)
			i--
		}
		p = l.dat0[i]

		i++
	}
}

// main lexing pass
func (l *Lexer) pass2() {
	for {
		if l.idx2 >= len(l.dat0) {
			break
		}

		r := l.dat0[l.idx2]
		switch r {
		case '\n':
			l.addToken(token.NLine, " ")
			l.line++
			l.char = 0
		case ' ', '\r', '\t':
			l.addToken(token.Space, " ")
		case ',':
			l.addToken(token.Comma, ",")
		case '.':
			l.addToken(token.Chain, ".")
		case '{':
			l.addToken(token.BraceL, "{")
		case '}':
			l.addToken(token.BraceR, "}")
		case '[':
			l.addToken(token.BrackL, "[")
		case ']':
			l.addToken(token.BrackR, "]")
		case '(':
			l.addToken(token.ParenL, "(")
		case ')':
			l.addToken(token.ParenR, ")")
		case ':':
			l.addToken(token.Typed, ":")
		case '=':
			l.addToken(token.Assign, "=")
		case '"', '\'':
			l.pass2_text()
		default:
			if unicode.IsDigit(r) {
				l.pass2_numb()
			} else if unicode.IsLetter(r) {
				l.pass2_word()
			} else if runtime.Opers.Contains(r) {
				l.pass2_oper()
			}
		}

		l.idx2++
		l.char++
	}

	l.idx2 = 0
}

// lex numeric value
func (l *Lexer) pass2_numb() {
	ranged := false
	decimal := false

	var builder strings.Builder

	for {
		if l.idx2 >= len(l.dat0) {
			break
		}

		r := l.dat0[l.idx2]

		if r == '.' && l.idx2+1 < len(l.dat0) && l.dat0[l.idx2+1] == '.' {
			ranged = true
			break
		}

		if (decimal && r == '.') || (r != '.' && !unicode.IsDigit(r)) {
			break
		}

		builder.WriteRune(r)

		if r == '.' {
			decimal = true
		}

		l.idx2++
		l.char++
	}

	l.idx2--
	l.char--
	l.addToken(token.Numb, builder.String())

	if ranged {
		l.idx2 += 2
		l.char += 2
		l.addToken(token.Range, "..")
	}
}

// lex text value
func (l *Lexer) pass2_text() {
	q := l.dat0[l.idx2]

	l.idx2++ // skip opening quote
	l.char++

	var builder strings.Builder

	builder.WriteRune(q)

	for {
		if l.idx2 >= len(l.dat0) {
			break
		}

		r := l.dat0[l.idx2]

		builder.WriteRune(r)

		if r == q {
			l.addToken(token.Text, builder.String())
			break
		}

		l.idx2++
		l.char++
	}

}

// lex keywords
func (l *Lexer) pass2_word() {
	var builder strings.Builder

	for {
		if l.idx2 >= len(l.dat0) {
			break
		}

		r := l.dat0[l.idx2]

		if r != '_' && !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			break
		}

		builder.WriteRune(r)

		l.idx2++
		l.char++
	}

	l.idx2--
	l.char++

	var tokenType = token.Symb
	var tokenData = builder.String()

	if tokenData == "true" || tokenData == "false" {
		tokenType = token.Bool
	} else if words.WordsAsStrings.Contains(tokenData) {
		tokenType = token.Word
	}

	l.addToken(tokenType, tokenData)
}

// lex operators
func (l *Lexer) pass2_oper() {
	r := l.dat0[l.idx2]

	if r == '/' && l.idx2+1 < len(l.dat0) && l.dat0[l.idx2+1] == '/' {
		for {
			if l.idx2 >= len(l.dat0) {
				break
			}

			r := l.dat0[l.idx2]
			if r == '\n' {
				break
			}

			l.idx2++
			l.char++
		}

		l.line++
		l.char = 0

		return
	}

	l.addToken(token.Oper, string(r))
}

// augmentation pass
func (l *Lexer) pass3() {
	l.pass3_decimal_and_operator_augmentation()
	l.pass3_negatives_and_extras_augmentation()
}

// augment period prefixed decimals and composed operators
func (l *Lexer) pass3_decimal_and_operator_augmentation() {
	l.data = token.ExecuteAugmentation(l.data, lexer_augmenters[0])
}

// augment negative numbers and anything else
func (l *Lexer) pass3_negatives_and_extras_augmentation() {
	l.data = token.ExecuteAugmentation(l.data, lexer_augmenters[1])
	/*for _, data := range l.data {
		fmt.Println(data.String())
	}
	fmt.Println("=============")*/
}

// cleanup pass (opt)
func (l *Lexer) pass4() {
	l.pass4_remove_extra_values(true, false, true)
}

// remove spaces and new lines optionally
func (l *Lexer) pass4_remove_extra_values(removeSpaces, removeNewLines, removeDuplicateSpaces bool) {
	if !removeSpaces && !removeNewLines && !removeDuplicateSpaces {
		return
	}

	i := 0
	p := l.data[i]
	for i < len(l.data) {
		data := l.data[i]

		if (removeSpaces && data.Type == token.Space) || (removeNewLines && data.Type == token.NLine) || (i != 0 && removeDuplicateSpaces && (p.Type == token.Space && data.Type == token.Space)) {
			l.data = append(l.data[:i], l.data[i+1:]...)
			i--
		}

		if i > 0 {
			p = l.data[i]
		}

		i++
	}
}
