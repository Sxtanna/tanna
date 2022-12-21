package compile

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joomcode/errorx"
	uuid "github.com/nu7hatch/gouuid"
	"golang.org/x/exp/slices"
	"tanna/commons/token"
	"tanna/commons/words"
	"tanna/compile/utils"
	"tanna/runtime"
)

const temp_replace_me = "5a2470d7"

var parser_augmenters = [][]*token.Augmenter{
	{
		{
			SkipCount: 1,
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Symb
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "+="
			},
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{
					{
						Type: here.Type,
						Data: here.Data,
					},
					{
						Type: token.Assign,
						Data: "=",
					},
					{
						Type: here.Type,
						Data: here.Data,
					},
					{
						Type: token.Oper,
						Data: "+",
					},
				}
			},
		},
		{
			SkipCount: 1,
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Symb
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "-="
			},
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{
					{
						Type: here.Type,
						Data: here.Data,
					},
					{
						Type: token.Assign,
						Data: "=",
					},
					{
						Type: here.Type,
						Data: here.Data,
					},
					{
						Type: token.Oper,
						Data: "-",
					},
				}
			},
		},
		{
			SkipCount: 1,
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Symb
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "/="
			},
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{
					{
						Type: here.Type,
						Data: here.Data,
					},
					{
						Type: token.Assign,
						Data: "=",
					},
					{
						Type: here.Type,
						Data: here.Data,
					},
					{
						Type: token.Oper,
						Data: "/",
					},
				}
			},
		},
		{
			SkipCount: 1,
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Symb
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "*="
			},
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{
					{
						Type: here.Type,
						Data: here.Data,
					},
					{
						Type: token.Assign,
						Data: "=",
					},
					{
						Type: here.Type,
						Data: here.Data,
					},
					{
						Type: token.Oper,
						Data: "*",
					},
				}
			},
		},
		{
			SkipCount: 1,
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Symb
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "++"
			},
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{
					{
						Type: token.ParenL,
						Data: "_",
					},
					{
						Type: here.Type,
						Data: here.Data,
					},
					{
						Type: token.Assign,
						Data: "=",
					},
					{
						Type: here.Type,
						Data: here.Data,
					},
					{
						Type: token.Oper,
						Data: "+",
					},
					{
						Type: token.Numb,
						Data: "1",
					},
					{
						Type: token.ParenR,
						Data: ")",
					},
					/*{
						Type: here.Type,
						Data: here.Data,
					},*/
				}
			},
		},
		{
			SkipCount: 1,
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Symb
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "--"
			},
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{
					{
						Type: token.ParenL,
						Data: "_",
					},
					{
						Type: here.Type,
						Data: here.Data,
					},
					{
						Type: token.Assign,
						Data: "=",
					},
					{
						Type: here.Type,
						Data: here.Data,
					},
					{
						Type: token.Oper,
						Data: "-",
					},
					{
						Type: token.Numb,
						Data: "1",
					},
					{
						Type: token.ParenR,
						Data: ")",
					},
					/*{
						Type: here.Type,
						Data: here.Data,
					},*/
				}
			},
		},
		{
			SkipCount: 1,
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.ParenL
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Word && data.Data == string(words.Sout)
			},
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{
					{
						Type: token.ParenL,
						Data: "_",
					},
					{
						Type: token.Word,
						Data: next.Data,
					},
				}
			},
		},
	},
	{
		{
			SkipCount: 1,
			HereMatch: func(data *token.Data) bool {
				return data.Type == token.Oper && data.Data == "++"
			},
			NextMatch: func(data *token.Data) bool {
				return data.Type == token.Symb
			},
			IntoToken: func(here, prev, next *token.Data) []*token.Data {
				return []*token.Data{
					{
						Type: token.ParenL,
						Data: "_",
					},
					{
						Type: next.Type,
						Data: next.Data,
					},
					{
						Type: token.Assign,
						Data: "=",
					},
					{
						Type: next.Type,
						Data: next.Data,
					},
					{
						Type: token.Oper,
						Data: "+",
					},
					{
						Type: token.Numb,
						Data: "1",
					},
					{
						Type: token.ParenR,
						Data: ")",
					},
					{
						Type: next.Type,
						Data: next.Data,
					},
				}
			},
		},
	},
}

type Parser struct {
	data []*token.Data
	cmds []runtime.Command
	idx1 int

	fail error

	parsingMain       int
	parsingWord       int
	parsingWhenExpr   int
	parsingLoopExpr   int
	parsingTuple      int
	parsingRange      int
	parsingProperty   int
	parsingFunction   int
	parsingParameters int
	parsingExpression int

	lastTempVarName string
}

func NewParser(data []*token.Data) *Parser {
	return &Parser{data: data}
}

// RunParser executes every pass of the parser at once
// p.pass0()
func (p *Parser) RunParser() {
	p.pass0()
	if !p.continues() {
		return
	}
	p.pass1()
	p.pass2()
}

func (p *Parser) ParserData() []runtime.Command {
	return p.cmds
}

func (p *Parser) Err() error {
	return p.fail
}

// continues whether the parser can continue, true if fail is nil and the cursor idx1 is not at the end of the data list
func (p *Parser) continues() bool {
	return p.fail == nil && p.idx1 < len(p.data) && p.ParsingDepth() < 100
}

func (p *Parser) current() *token.Data {
	return p.data[p.idx1]
}

func (p *Parser) ParsingDepth() int {
	return p.parsingMain + p.parsingWord + p.parsingWhenExpr + p.parsingLoopExpr + p.parsingTuple + p.parsingProperty + p.parsingFunction + p.parsingParameters + p.parsingExpression
}

// pass0 executes the parser's augmenters; these expand certain token patterns, ex.
//
//	a++ -> a = a + 1;
func (p *Parser) pass0() {
	p.data = token.ExecuteAugmentation(p.data, parser_augmenters[0])
	p.data = token.ExecuteAugmentation(p.data, parser_augmenters[1])

	for _, data := range p.data {
		fmt.Println(data.String())
	}
}

// pass1 executes the main parsing functions, separating each block of parsing with a command.CommandBoundary
func (p *Parser) pass1() {
	for {
		if !p.continues() {
			break
		}

		if len(p.cmds) > 0 {
			if _, ok := p.cmds[len(p.cmds)-1].(*runtime.CommandBoundary); !ok {
				p.cmds = append(p.cmds, &runtime.CommandBoundary{})
			}
		}

		if err := p._main(&p.cmds); err != nil {
			p.fail = err
			return
		}
	}

	p.idx1 = 0
}

// pass2 basic cleanup pass to remove trailing command boundary instances
func (p *Parser) pass2() {
	if len(p.cmds) == 0 {
		return
	}

	for {
		last := p.cmds[len(p.cmds)-1]

		if _, ok := last.(*runtime.CommandBoundary); !ok {
			break
		} else {
			p.cmds = p.cmds[:len(p.cmds)-1]
		}
	}
}

func (p *Parser) _main(cmds *[]runtime.Command) (err error) {
	p.parsingMain++
	defer func() { p.parsingMain-- }()

	p._skip_new_lines()

	if !p.continues() {
		return
	}

	t := p.data[p.idx1]

	switch t.Type {
	case token.Space:
		fallthrough
	case token.NLine:
		p.idx1++ // skip spaces and new lines
		return
	case token.Word:
		return p._word(cmds)
	case token.ParenL:
		if expr, e := p._eval_expr("main"); e != nil {
			return e
		} else {
			*cmds = append(*cmds, expr...)
		}
		return
	case token.Symb:
		return p._symbol(cmds)
	default:
		return TokenTypeOutOfPlace.
			New("could not parse token %v: `%v` is not handled in main context", t, t.Type)
	}
}

func (p *Parser) _eval(cmds []runtime.Command) (out []runtime.Command, err error) {
	if len(cmds) <= 1 {
		return cmds, nil
	}

	stack := runtime.NewStack()
	shunt := make([]runtime.Command, 0)

	for _, cmd := range cmds {
		if cmdOperate, ok := cmd.(*runtime.CommandOperator); !ok {
			shunt = append(shunt, cmd)
			continue
		} else {
			switch cmdOperate.Oper {
			case runtime.OperatorSOS:
				stack.Push(cmd)
			case runtime.OperatorEOS:
				for stack.Len() != 0 && stack.Peek() != nil && (stack.Peek().(*runtime.CommandOperator)).Oper != runtime.OperatorSOS {
					if value, ok := stack.Pull(); ok {
						shunt = append(shunt, value.(runtime.Command))
					}
				}

				if _, ok := stack.Pull(); !ok {
					return nil, ShuntingYardEmptyStack.New("stack is empty after flushing for EOS")
				}
			default:
				for stack.Len() != 0 {
					thisOper := cmdOperate.Oper
					nextOper := stack.Peek().(*runtime.CommandOperator).Oper

					if nextOper == runtime.OperatorSOS {
						break
					}

					if nextOper.Precedence <= thisOper.Precedence && !(nextOper.Precedence == thisOper.Precedence && thisOper.Associativity == runtime.LTR) {
						break
					}

					if value, ok := stack.Pull(); !ok {
						return nil, ShuntingYardEmptyStack.New("stack is empty during default")
					} else {
						shunt = append(shunt, value.(runtime.Command))
					}
				}

				stack.Push(cmd)
			}
		}
	}

	for {
		if value, ok := stack.Pull(); !ok {
			break
		} else {
			shunt = append(shunt, value.(runtime.Command))
		}
	}

	return shunt, nil
}

func (p *Parser) _expr() (cmds []runtime.Command, err error) {
	p.parsingExpression++
	defer func() { p.parsingExpression-- }()

	cmds = make([]runtime.Command, 0)

	openParen := 0
	headIndex := p.idx1

	for {
		if !p.continues() {
			return
		}

		t := p.data[p.idx1]

		if t.Type == token.Comma && p.parsingTuple > 0 {
			return cmds, nil
		}

		if t.Type == token.Range && p.parsingRange > 0 {
			return cmds, nil
		}

		if t.Type == token.BraceL && (p.parsingWhenExpr > 0 || p.parsingLoopExpr > 0) {
			return cmds, nil
		}

		switch t.Type {
		case token.NLine:
			return
		case token.Bool:
			fallthrough
		case token.Numb:
			fallthrough
		case token.Text:
			err = p._literal(&cmds)
		case token.Symb:
			err = p._symbol(&cmds)
		case token.ParenL:
			p.idx1++

			openParen++
			cmds = append(cmds, &runtime.CommandOperator{Oper: runtime.OperatorSOS})
		case token.ParenR:
			if openParen == 0 {
				return cmds, nil
			}
			p.idx1++

			openParen--
			cmds = append(cmds, &runtime.CommandOperator{Oper: runtime.OperatorEOS})
		case token.Oper:
			if oper, err := p._operator(); err != nil {
				return nil, errorx.Decorate(err, "could not parse operator in expression")
			} else {
				cmds = append(cmds, &runtime.CommandOperator{Oper: oper})
			}
		case token.Comma:
			if openParen < 1 {
				return nil, TokenTypeOutOfPlace.
					New("could not parse token %v: `%v` commas in expressions require parens", t, t.Type)
			}

			backtrack := 0
			for p.data[p.idx1].Type != token.ParenL || p.idx1 > headIndex {
				p.idx1--
				backtrack++
			}

			for backtrack > 0 && len(cmds) > 0 {
				backtrack--
				cmds = cmds[:len(cmds)-1]
			}

			if tuple, e := p._tuple(); e != nil {
				return nil, errorx.Decorate(err, "could not parse tuple in expression")
			} else {
				cmds = append(cmds, &runtime.CommandRoute{Route: runtime.NewRoute(utils.FlattenSlices(tuple))})
				cmds = append(cmds, &runtime.CommandTuple{Size: len(tuple)})
			}
		case token.Word:
			if word, e := p._word_find(t); e != nil {
				return nil, e
			} else {
				if word == words.Push || word == words.Pull || word == words.Sout || word == words.Range || word == words.By {
					err = p._word(&cmds)
				} else {
					return nil, TokenDataOutOfPlace.
						New("could not parse token %v: `%s` is not a valid word for use in expressions", t, t.Data)
				}
			}
		default:
			return nil, TokenTypeOutOfPlace.
				New("could not parse token %v: `%v` is not handled in expr context", t, t.Type)
		}

		if err != nil {
			return nil, err
		}
	}
}

func (p *Parser) _eval_expr(name string) (out []runtime.Command, err error) {
	if expr, e := p._expr(); e != nil {
		return nil, errorx.Decorate(e, fmt.Sprintf("could not parse expression [%s]", name))
	} else if yard, e := p._eval(expr); e != nil {
		return nil, errorx.Decorate(e, fmt.Sprintf("could not apply shunting yard to expression [%s]", name))
	} else {
		return yard, nil
	}
}

func (p *Parser) _word(cmds *[]runtime.Command) (err error) {
	p.parsingWord++
	defer func() { p.parsingWord-- }()

	if !p.continues() {
		return
	}

	w := p.data[p.idx1]
	if w.Type != token.Word {
		return TokenTypeOutOfPlace.
			New("could not parse token %v: `%v` is not a word token type", w, w.Type)
	}

	if word, e := p._word_find(w); e != nil {
		return e
	} else {
		p.idx1++ // skip word token

		switch word {
		case words.Variable:
			fallthrough
		case words.Constant:
			return p._property(cmds, word == words.Constant)
		case words.Function:
			return p._function(cmds)
		case words.Loop:
			return p._loop(cmds)
		case words.When:
			return p._when(cmds)
		case words.Sout:
			returnsToStack := false

			if p.continues() && p.data[p.idx1].Type == token.Return {
				p.idx1++ // skip return token

				returnsToStack = true
			}

			if expr, e := p._eval_expr("sout"); e != nil {
				return e
			} else {
				*cmds = append(*cmds, &runtime.CommandRoute{Route: runtime.NewRoute(expr)})
				*cmds = append(*cmds, &runtime.CommandSout{
					AppendNewLine: true,
					ReturnToStack: returnsToStack,
				})
			}
		case words.Push:
			if expr, e := p._eval_expr("push"); e != nil {
				return e
			} else {
				*cmds = append(*cmds, &runtime.CommandRoute{
					Route: runtime.NewRoute(expr),
				})
			}
		case words.Pull:
			*cmds = append(*cmds, &runtime.CommandPull{
				PullValueFromStack: p.parsingExpression <= 0,
			})
		case words.Stop:
			*cmds = append(*cmds, &runtime.CommandLoopStop{FullStop: true})
		case words.Cont:
			*cmds = append(*cmds, &runtime.CommandLoopStop{FullStop: false})
		case words.Range:
			return p._range(cmds)
		case words.By:
			if expr, e := p._eval_expr("by"); e != nil {
				return e
			} else {
				*cmds = append(*cmds, &runtime.CommandRoute{
					Route: runtime.NewRoute(expr),
				})
				*cmds = append(*cmds, &runtime.CommandMakeRangeIter{})
			}
		default:
			return TokenDataOutOfPlace.
				New("could not parse token %v: `%v` word is not handled in word context", w, w.Data)
		}
	}

	return
}

func (p *Parser) _word_find(token *token.Data) (word words.Word, err error) {
	if !words.WordsAsStrings.Contains(token.Data) {
		return "", TokenDataOutOfPlace.
			New("could not parse token %v: `%s` is not a valid word", token, token.Data)
	}

	return words.Word(token.Data), nil
}

func (p *Parser) _property(cmds *[]runtime.Command, constant bool) (err error) {
	p.parsingProperty++
	defer func() { p.parsingProperty-- }()

	if !p.continues() {
		return
	}

	s := p.data[p.idx1]
	if s.Type != token.Symb {
		return TokenTypeOutOfPlace.
			New("could not parse token %v: `%v` is not a symb token type", s, s.Type)
	}

	p.idx1++ // skip symb token

	property := runtime.NewProperty(s.Data, constant)

	if p.continues() && p.data[p.idx1].Type == token.Typed {
		p.idx1++ // skip type specifier token

		if model, e := p._type(); e != nil {
			return e
		} else {
			property.Model = model
		}
	}

	if !p.continues() {
		return TokenRequiredMissing.
			New("could not parse property: expected to find `%v` token type", token.Assign)
	}

	a := p.data[p.idx1]
	if a.Type != token.Assign {
		return TokenTypeOutOfPlace.
			New("could not parse token %v: `%v` is not a assign token type", a, a.Type)
	}

	p.idx1++ // skip assign token

	*cmds = append(*cmds, &runtime.CommandPropertyDefine{Property: property})

	if expr, e := p._eval_expr("property assignment"); e != nil {
		return e
	} else {
		*cmds = append(*cmds, expr...)
		*cmds = append(*cmds, &runtime.CommandPropertyAssign{PropertyName: property.Name})
	}

	return
}

func (p *Parser) _function(cmds *[]runtime.Command) (err error) {
	p.parsingFunction++
	defer func() { p.parsingFunction-- }()

	if !p.continues() {
		return
	}

	s := p.data[p.idx1]
	if s.Type != token.Symb {
		return TokenTypeOutOfPlace.
			New("could not parse token %v: `%v` is not a symb token type", s, s.Type)
	}

	p.idx1++ // skip symb token

	function := runtime.NewFunction(s.Data)

	names := make([]string, 0)

	if !p.continues() {
		return TokenRequiredMissing.
			New("could not parse function: expected to find one of [`%v`, `%v`,`%v`] token type", token.ParenL, token.Typed, token.BraceL)
	}

	if p.data[p.idx1].Type == token.ParenL {
		p.idx1++ // skip param parens

		group := make([]string, 0)

		for {
			if !p.continues() {
				break
			}

			t := p.data[p.idx1]

			if t.Type == token.Comma {
				if len(group) == 0 && len(function.Accepts) == 0 {
					return TokenTypeOutOfPlace.
						New("could not parse token %v: `%v` first value in function parens must be a parameter", t, t.Type)
				}

				p.idx1++ // skip comma token

				continue
			} else if t.Type == token.Symb {
				if slices.Contains(names, t.Data) {
					return TokenDataOutOfPlace.
						New("could not parse token %v: `%v` function parameter with name %s already exists", t, t.Type, t.Data)
				}

				p.idx1++ // skip symb token

				names = append(names, t.Data)
				group = append(group, t.Data)

				continue
			} else if t.Type == token.Typed {
				p.idx1++ // skip typed token

				if model, e := p._type(); e != nil {
					return errorx.Decorate(e, "could not parse function parameter type")
				} else {
					for _, name := range group {
						function.AddAccepts(name, model)
					}
				}

				group = nil // clear name group

				continue
			} else if t.Type == token.ParenR {
				if len(group) != 0 {
					return TokenTypeOutOfPlace.
						New("could not parse token %v: `%v` unfinished parameter group %s", t, t.Type, group)
				}

				p.idx1++ // skip r paren token

				break
			} else {
				return TokenTypeOutOfPlace.
					New("could not parse token %v: `%v` function parameters does not accept this", t, t.Type)
			}
		}
	}

	if !p.continues() {
		return TokenRequiredMissing.
			New("could not parse function: expected to find one of [`%v`, `%v`] token type", token.Typed, token.BraceL)
	}

	if p.data[p.idx1].Type == token.Typed {
		p.idx1++ // skip typed parens

		if model, e := p._type(); e != nil {
			return errorx.Decorate(e, "could not parse function return type")
		} else {
			switch m := model.(type) {
			case *runtime.ModelTanna:
				function.AddReturns("$0", m)
			case *runtime.ModelPlatform:
				function.AddReturns("$0", m)
			case *runtime.ModelTuple:
				for i, part := range m.Parts() {
					function.AddReturns(fmt.Sprintf("$%d", i), part)
				}
			}
		}
	}

	if !p.continues() {
		return TokenRequiredMissing.
			New("could not parse function: expected to find one of [`%v`, `%v`] token type", token.Return, token.BraceL)
	}

	if p.data[p.idx1].Type != token.Return && p.data[p.idx1].Type != token.BraceL {
		return TokenTypeOutOfPlace.
			New("could not parse token %v: `%v` function expression or body requires one of [`%v`, `%v`]", p.data[p.idx1], p.data[p.idx1].Type, token.Return, token.BraceL)
	}

	expression := p.data[p.idx1].Type == token.Return

	p.idx1++ // skip return or l brace token

	body := make([]runtime.Command, 0)

	for _, parameter := range function.Accepts {
		property := runtime.NewProperty(parameter.Name, true)
		property.Model = parameter.Model

		body = append(body, &runtime.CommandPropertyDefine{Property: property})
		body = append(body, &runtime.CommandPropertyAssign{PropertyName: property.Name})
	}

	if expression {
		if expr, e := p._eval_expr("function expression"); e != nil {
			return e
		} else {
			body = append(body, &runtime.CommandRoute{Route: runtime.NewRoute(expr)})
		}
	} else {
		for {
			if !p.continues() {
				break
			}

			t := p.data[p.idx1]

			if t.Type == token.BraceR {
				break
			}

			if t.Type == token.NLine {
				p.idx1++ // skip new line token
				continue
			}

			if t.Type == token.Return {
				p.idx1++ // skip return token

				if expr, e := p._eval_expr("function body return"); e != nil {
					return e
				} else {
					body = append(body, &runtime.CommandRoute{Route: runtime.NewRoute(expr)})
				}

				break
			}

			if e := p._main(&body); e == nil {
				continue
			} else {
				return errorx.Decorate(e, "could not parse function body")
			}

			// return errors.TokenTypeOutOfPlace.
			// 	New("could not parse token %v: `%v` function body does not accept this", t, t.Type)
		}

		p._skip_new_lines()

		if !p.continues() {
			return TokenRequiredMissing.
				New("could not parse function: expected to find `%v` token type", token.BraceR)
		}

		p.idx1++ // skip r brace token
	}

	for _, parameter := range function.Accepts {
		body = append(body, &runtime.CommandPropertyNilify{PropertyName: parameter.Name})
	}

	function.Body = runtime.NewRoute(body)

	*cmds = append(*cmds, &runtime.CommandFunctionDefine{Function: function})

	return
}

func (p *Parser) _literal(cmds *[]runtime.Command) (err error) {
	if !p.continues() {
		return
	}

	l := p.data[p.idx1]

	switch l.Type {
	case token.Bool:
		var value bool

		if l.Data == "true" {
			value = true
		} else if l.Data == "false" {
			value = false
		} else {
			return TokenDataOutOfPlace.
				New("could not parse token %v: `%v` is not a valid bit value", l, l.Data)
		}

		*cmds = append(*cmds, &runtime.CommandLiteral{Value: &runtime.Value{
			Model: runtime.FindBuiltInModel("Bit"),
			Value: value,
		}})
	case token.Numb:
		if !strings.ContainsRune(l.Data, '.') {
			if i, err := strconv.ParseInt(l.Data, 10, 64); err != nil {
				return errorx.Decorate(err, "could not parse Int literal %v", l.Data)
			} else {
				*cmds = append(*cmds, &runtime.CommandLiteral{Value: &runtime.Value{
					Model: runtime.FindBuiltInModel("Int"),
					Value: i,
				}})
			}
		} else {
			if f, err := strconv.ParseFloat(l.Data, 64); err != nil {
				return errorx.Decorate(err, "could not parse Dec literal %v", l.Data)
			} else {
				*cmds = append(*cmds, &runtime.CommandLiteral{Value: &runtime.Value{
					Model: runtime.FindBuiltInModel("Dec"),
					Value: f,
				}})
			}
		}
	case token.Text:
		if strings.HasPrefix(l.Data, `'`) {
			if len(l.Data) > 1 {
				return TokenDataOutOfPlace.
					New("could not parse token %v: `%v` is not a valid let value", l, l.Data)
			}

			*cmds = append(*cmds, &runtime.CommandLiteral{Value: &runtime.Value{
				Model: runtime.FindBuiltInModel("Let"),
				Value: strings.Trim(l.Data, `'`),
			}})
		} else if strings.HasPrefix(l.Data, `"`) {
			*cmds = append(*cmds, &runtime.CommandLiteral{Value: &runtime.Value{
				Model: runtime.FindBuiltInModel("Txt"),
				Value: strings.Trim(l.Data, `"`),
			}})
		} else {
			return TokenDataOutOfPlace.
				New("could not parse token %v: `%v` is not a valid text literal", l, l.Data)
		}
	default:
		return TokenTypeOutOfPlace.
			New("could not parse token %v: `%v` is not a valid literal token", l, l.Type)
	}

	p.idx1++

	return
}

func (p *Parser) _range(cmds *[]runtime.Command) (err error) {
	p.parsingRange++
	defer func() { p.parsingRange-- }()

	if !p.continues() || p.current().Type != token.ParenL {
		return TokenRequiredMissing.
			New("could not parse range: expected to find `%v` token type", token.ParenL)
	}

	p.idx1++ // skip l paren token

	min := make([]runtime.Command, 0)
	max := make([]runtime.Command, 0)

	if out, e := p._eval_expr("min range expression"); e != nil {
		return errorx.Decorate(e, "could not parse min range expression")
	} else {
		min = append(min, out...)
	}

	if p.continues() && p.current().Type == token.Range {
		p.idx1++ // skip range token
	} else {
		return TokenRequiredMissing.
			New("could not parse range: expected to find `%v` token type", token.Range)
	}

	if out, e := p._eval_expr("max range expression"); e != nil {
		return errorx.Decorate(e, "could not parse max range expression")
	} else {
		max = append(max, out...)
	}

	if !p.continues() || p.current().Type != token.ParenR {
		return TokenRequiredMissing.
			New("could not parse range: expected to find `%v` token type closing the expression, found '%v'", token.ParenR)
	}

	p.idx1++ // skip r paren token

	*cmds = append(*cmds, &runtime.CommandRoute{Route: runtime.NewRoute(min)})
	*cmds = append(*cmds, &runtime.CommandRoute{Route: runtime.NewRoute(max)})

	*cmds = append(*cmds, &runtime.CommandMakeRange{Num: true})

	return
}

func (p *Parser) _tuple() (out [][]runtime.Command, err error) {
	p.parsingTuple++
	defer func() { p.parsingTuple-- }()

	if !p.continues() {
		return
	}

	l := p.data[p.idx1]
	if l.Type != token.ParenL {
		return nil, TokenTypeOutOfPlace.
			New("could not parse token %v: `%v` tuple must start with left paren", l, l.Type)
	}

	p.idx1++ // skip l paren token

	out = make([][]runtime.Command, 0)

	for {
		if !p.continues() {
			return
		}

		t := p.data[p.idx1]

		if t.Type == token.Comma {
			if len(out) == 0 {
				return nil, TokenTypeOutOfPlace.
					New("could not parse token %v: `%v` first tuple value required before a comma", t, t.Type)
			}

			p.idx1++ // skip comma token

			continue
		} else if t.Type == token.ParenR {
			if len(out) == 0 && p.parsingParameters <= 0 {
				return nil, TokenTypeOutOfPlace.
					New("could not parse token %v: `%v` tuples must contain data", t, t.Type)
			}

			break
		} else if t.Type == token.ParenL {
			marker := p.data[p.idx1]

			if nested, e := p._tuple(); e != nil {
				return nil, errorx.Decorate(e, "could not parse nested tuple")
			} else {
				values := utils.FlattenSlices(nested)

				if marker.Data != "_" {
					values = append(values, &runtime.CommandTuple{Size: len(nested)})
				}

				out = append(out, values)
			}

			continue
		}

		if expr, e := p._eval_expr("tuple element"); e != nil {
			return nil, e
		} else {
			out = append(out, expr)
		}
	}

	r := p.data[p.idx1]
	if r.Type != token.ParenR {
		return nil, TokenTypeOutOfPlace.
			New("could not parse token %v: `%v` tuple must end with right paren", r, r.Type)
	}

	p.idx1++ // skip r paren token

	return out, nil
}

func (p *Parser) _symbol(cmds *[]runtime.Command) (err error) {
	if !p.continues() {
		return
	}

	s := p.data[p.idx1]

	p.idx1++ // skip over symbol token

	if !p.continues() {
		return p._symbol_call_property(cmds, s.Data)
	}

	n := p.data[p.idx1]

	switch n.Type {
	case token.ParenL:
		return p._symbol_call_function(cmds, s.Data)
	case token.Point:

	case token.Assign:
		p.idx1++ // skip assign token

		if expr, e := p._eval_expr("symbol assignment"); e != nil {
			return e
		} else {
			*cmds = append(*cmds, &runtime.CommandPropertyAccess{PropertyName: s.Data})
			*cmds = append(*cmds, &runtime.CommandRoute{Route: runtime.NewRoute(expr)})
			*cmds = append(*cmds, &runtime.CommandPropertyAssign{PropertyName: s.Data})
		}
	default:
		return p._symbol_call_property(cmds, s.Data)
	}

	return
}

func (p *Parser) _symbol_call_function(cmds *[]runtime.Command, name string) (err error) {
	p.parsingParameters++

	if out, e := p._tuple(); e != nil {
		return errorx.Decorate(e, "could not parse function parameters")
	} else {
		p.parsingParameters--

		for _, element := range out {
			*cmds = append(*cmds, &runtime.CommandRoute{Route: runtime.NewRoute(element)})
		}

		*cmds = append(*cmds, &runtime.CommandFunctionAccess{FunctionName: name})
	}

	return
}

func (p *Parser) _symbol_call_property(cmds *[]runtime.Command, name string) (err error) {

	/*if name == temp_replace_me {
		if p.lastTempVarName != "" {
			name = p.lastTempVarName
		} else if n, e := p._symbol_gen_temporary(cmds); e != nil {
			return e
		} else {
			name = n
		}
	}*/

	*cmds = append(*cmds, &runtime.CommandPropertyAccess{PropertyName: name})
	return
}

func (p *Parser) _symbol_gen_temporary(cmds *[]runtime.Command) (name string, err error) {
	if id, e := uuid.NewV4(); e != nil {
		return "", errorx.Decorate(e, "could not generate temporary variable uuid")
	} else {
		name = fmt.Sprintf("temp_%s", id.String())

		*cmds = append(*cmds, &runtime.CommandPropertyDefine{Property: runtime.NewProperty(name, true)})
		*cmds = append(*cmds, &runtime.CommandPropertyAssign{PropertyName: name})
		*cmds = append(*cmds, &runtime.CommandPropertyAccess{PropertyName: name})

		p.lastTempVarName = name

		return
	}
}

func (p *Parser) _operator() (oper *runtime.Operator, err error) {
	if !p.continues() {
		return
	}

	o := p.data[p.idx1]

	switch o.Data {
	case "+":
		oper = runtime.OperatorAdd
	case "-":
		oper = runtime.OperatorSub
	case "*":
		oper = runtime.OperatorMul
	case "/":
		oper = runtime.OperatorDiv
	case "!":
		oper = runtime.OperatorNot
	case "==":
		oper = runtime.OperatorSame
	case "!=":
		oper = runtime.OperatorDiff
	case "||":
		oper = runtime.OperatorElse
	case "&&":
		oper = runtime.OperatorBoth
	default:
		return nil, TokenTypeOutOfPlace.
			New("could not parse token %v: `%v` is not a valid operator token", o, o.Type)
	}

	p.idx1++ // move past operator

	return
}

func (p *Parser) _type() (model runtime.Model, err error) {
	if !p.continues() {
		return
	}

	t := p.data[p.idx1]
	switch t.Type {
	case token.Symb:
		return p._type_basic()
	case token.ParenL:
		return p._type_tuple()
	default:
		return nil, TokenTypeOutOfPlace.
			New("could not parse token %v: `%v` is not a valid type specifier", t, t.Type)
	}
}

func (p *Parser) _type_basic() (model runtime.Model, err error) {
	if !p.continues() {
		return
	}

	s := p.data[p.idx1]
	if s.Type != token.Symb {
		return nil, TokenTypeOutOfPlace.
			New("could not parse token %v: `%v` is not a valid basic type specifier", s, s.Type)
	}

	var basic runtime.Model

	if found := runtime.FindBuiltInModel(s.Data); found != nil {
		basic = found
	} else {
		basic = runtime.NewModelTanna(s.Data)
	}

	p.idx1++ // skip type symbol

	if p.continues() && p.data[p.idx1].Type == token.Bound { // shorthand tuple syntax
		p.idx1++ // skip shorthand bound

		if !p.continues() {
			return nil, TokenRequiredMissing.
				New("could not parse shorthand tuple: expected to find `%v` token type", token.Bound)
		} else {
			n := p.data[p.idx1]

			if n.Type != token.Numb {
				return nil, TokenTypeOutOfPlace.
					New("could not parse token %v: `%v` is not a tuple shorthand size specifier", n, n.Type)
			}

			if strings.ContainsRune(n.Data, '.') {
				return nil, TokenDataOutOfPlace.
					New("could not parse token %v: `%s` as a tuple shorthand size specifier should be a whole number", n, n.Data)
			}

			p.idx1++ // skip shorthand size specifier

			if size, e := strconv.Atoi(n.Data); e != nil {
				return nil, errorx.Decorate(e, "could not parse tuple shorthand size specifier")
			} else {
				models := make([]runtime.Model, size)

				for i := 0; i < size; i++ {
					models[i] = basic
				}

				return runtime.NewModelTuple(models), nil
			}
		}
	}

	model = basic

	return
}

func (p *Parser) _type_tuple() (model runtime.Model, err error) {
	if !p.continues() {
		return
	}

	return
}

func (p *Parser) _when(cmds *[]runtime.Command) (err error) {
	expr := make([]runtime.Command, 0)

	p.parsingWhenExpr++
	if e := p._when_expr(&expr); e != nil {
		return errorx.Decorate(e, "could not parse when expression")
	}
	p.parsingWhenExpr--

	pass := make([]runtime.Command, 0)

	if e := p._when_body(&pass); e != nil {
		return errorx.Decorate(e, "could not parse when pass body")
	}

	p._skip_new_lines()

	if !p.continues() || p.data[p.idx1].Type != token.BraceR {
		return TokenRequiredMissing.
			New("could not parse when pass body: expected to find `%v` token type", token.BraceR)
	}

	p.idx1++ // skip r brace token

	p._skip_new_lines()

	when := &runtime.CommandWhen{
		Expr: runtime.NewRoute(expr),
		Pass: runtime.NewRoute(pass),
	}

	if p.continues() && p.data[p.idx1].Type == token.Word {
		if word, e := p._word_find(p.data[p.idx1]); e != nil {
			return e
		} else {
			if word == words.Else {
				p.idx1++ // skip else token

				if !p.continues() {
					return TokenRequiredMissing.
						New("could not parse when: expected to find `%v` token type", token.BraceL)
				}

				if p.data[p.idx1].Type == token.BraceL {
					p.idx1++ // skip l brace token
				} else {
					return TokenTypeOutOfPlace.
						New("could not parse token %v: `%v` when body requires a left brace", p.data[p.idx1], p.data[p.idx1].Type)
				}

				fail := make([]runtime.Command, 0)

				if e := p._when_body(&fail); e != nil {
					return errorx.Decorate(e, "could not parse when fail body")
				}

				p._skip_new_lines()

				if !p.continues() || p.data[p.idx1].Type != token.BraceR {
					return TokenRequiredMissing.
						New("could not parse when fail body: expected to find `%v` token type", token.BraceR)
				}

				p.idx1++ // skip r brace token

				p._skip_new_lines()

				when.Fail = runtime.NewRoute(fail)
			}
		}
	}

	*cmds = append(*cmds, when)

	return
}

func (p *Parser) _when_body(cmds *[]runtime.Command) (err error) {
	for {
		if !p.continues() {
			return
		}

		t := p.data[p.idx1]

		if t.Type == token.BraceR {
			return
		}

		if t.Type == token.NLine {
			p._skip_new_lines()
			continue
		}

		if e := p._main(cmds); e == nil {
			continue
		} else {
			return errorx.Decorate(e, "could not parse main when body")
		}
	}
}

func (p *Parser) _when_expr(cmds *[]runtime.Command) (err error) {

	if !p.continues() {
		return TokenRequiredMissing.
			New("could not parse when: expected to find one of [`%v`, `%v`] token type", token.ParenL, token.BraceL)
	}

	if p.data[p.idx1].Type != token.BraceL {
		insideParentheses := p.data[p.idx1].Type == token.ParenL && p.data[p.idx1].Data == "("

		if insideParentheses {
			p.idx1++ // skip l paren token
		}

		if out, e := p._eval_expr("when expression"); e != nil {
			return errorx.Decorate(e, "could not parse when expression")
		} else {
			*cmds = append(*cmds, out...)
		}

		if insideParentheses {
			if p.data[p.idx1].Type == token.ParenR {
				p.idx1++ // skip r paren token
			} else {
				return TokenRequiredMissing.
					New("could not parse when: expected to find `%v` token type closing the expression", token.ParenR)
			}
		}
	}

	if !p.continues() {
		return TokenRequiredMissing.
			New("could not parse when: expected to find `%v` token type", token.BraceL)
	}

	if p.data[p.idx1].Type == token.BraceL {
		p.idx1++ // skip l brace token
	} else {
		return TokenTypeOutOfPlace.
			New("could not parse token %v: `%v` when expr requires a left brace", p.data[p.idx1], p.data[p.idx1].Type)
	}

	return
}

func (p *Parser) _loop(cmds *[]runtime.Command) (err error) {
	loop := &runtime.CommandLoop{}

	p.parsingLoopExpr++
	if e := p._loop_expr(cmds); e != nil {
		return errorx.Decorate(e, "could not parse loop expression")
	} else {
		loop.Expr = runtime.LoopExpressionDeferred
	}
	p.parsingLoopExpr--

	body := make([]runtime.Command, 0)

	if e := p._loop_body(&body); e != nil {
		return errorx.Decorate(e, "could not parse loop body")
	}

	loop.Body = runtime.NewRoute(body)

	p._skip_new_lines()

	if !p.continues() || p.data[p.idx1].Type != token.BraceR {
		return TokenRequiredMissing.
			New("could not parse loop: expected to find `%v` token type", token.BraceR)
	}

	p.idx1++ // skip r brace token

	*cmds = append(*cmds, loop)

	return
}

func (p *Parser) _loop_body(cmds *[]runtime.Command) (err error) {
	for {
		if !p.continues() {
			return
		}

		t := p.data[p.idx1]

		if t.Type == token.BraceR {
			return
		}

		if t.Type == token.NLine {
			p._skip_new_lines()
			continue
		}

		if e := p._main(cmds); e == nil {
			continue
		} else {
			return errorx.Decorate(e, "could not parse main loop body")
		}
	}
}

func (p *Parser) _loop_expr(cmds *[]runtime.Command) (err error) {
	if !p.continues() {
		return TokenRequiredMissing.
			New("could not parse loop: expected to find one of [`%v`, `%v`] token type", token.ParenL, token.BraceL)
	}

	if p.data[p.idx1].Type != token.BraceL {
		insideParentheses := p.data[p.idx1].Type == token.ParenL && p.data[p.idx1].Data == "("

		if insideParentheses {
			p.idx1++ // skip l paren token
		}

		if out, e := p._eval_expr("loop expression"); e != nil {
			return errorx.Decorate(e, "could not loop when expression")
		} else {
			*cmds = append(*cmds, &runtime.CommandRoute{Route: runtime.NewRoute(out)})
			// expr = &runtime.LoopExpressionBit{Expr: runtime.NewRoute(out)}
		}

		/*if p.continues() && p.current().Type == token.Symb {
			s := p.current()

			p.idx1++ // skip symbol

			if p.continues() && p.current().Type == token.Word && words.Word(p.current().Data) == words.In {
				p.idx1++ // skip in token

				vars := make([]runtime.Command, 0)
				vars = append(vars, &runtime.CommandPropertyDefine{Property: runtime.NewProperty(s.Data, false)})
				loop.Vars = runtime.NewRoute(vars)

				if p.continues() {

					switch p.current().Type {
					case token.Symb:

						call := make([]runtime.Command, 0)
						if e := p._symbol(&call); e != nil {
							return nil, e
						}


						expr = &runtime.LoopExpressionRange{}
					}

					if p.current().Type != token.Word || words.Word(p.current().Data) != words.Range {
						p.idx1++ // skip result


					} else {
						p.idx1++ // skip range token

						cmds := make([]runtime.Command, 0)

						if e := p._range(&cmds); e != nil {
							return nil, e
						}
					}
				}
			}
		} else {
			if out, e := p._eval_expr("loop expression"); e != nil {
				return nil, errorx.Decorate(e, "could not loop when expression")
			} else {
				expr = &runtime.LoopExpressionBit{Expr: runtime.NewRoute(out)}
			}
		}*/

		if insideParentheses {
			if p.data[p.idx1].Type == token.ParenR {
				p.idx1++ // skip r paren token
			} else {
				return MakeTokenRequiredMissing("loop", token.ParenR, p.data[p.idx1], "closing the expression")
			}
		}
	}

	if !p.continues() {
		return TokenRequiredMissing.
			New("could not parse loop: expected to find `%v` token type", token.BraceL)
	}

	if p.data[p.idx1].Type == token.BraceL {
		p.idx1++ // skip l brace token
	} else {
		return TokenTypeOutOfPlace.
			New("could not parse token %v: `%v` loop body requires a left brace", p.data[p.idx1], p.data[p.idx1].Type)
	}

	return
}

func (p *Parser) _skip_new_lines() {
	for p.continues() && p.data[p.idx1].Type == token.NLine {
		p.idx1++ // skip new lines
	}
}
