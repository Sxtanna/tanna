package compile

import (
	"log"
	"testing"

	"tanna/runtime"
)

const (
	parser_code_sout = `
sout 1
sout "Hello World!"
`
	parser_code_property = `
	constant t: Int::2 = 0
	// constant a: Int = 10
	// variable b: Int = 20 + 1
	// variable c = 26

	// constant bool = true
`
	parser_code_function = `
	function add(a: Int, b: Int): Int {
		=> a + b
	}

	add(value0, value1)
`
	parser_code_control_flow_when = `
	when (10 > 5) {
		sout "10 is greater than 5"
	} else (2 < 4) {
		sout "2 is less than 4"
	} else {
		sout "this definitely doesn't make sense..."
	}
`
	parser_code_control_flow_case = `
	case (20) {
		10 {
			sout "nope not here!!"
		}
		20 {
			sout "definitely this one"
		}
	}
`
	parser_code_control_flow_loop = `
	loop (i 0..10) {
		when (i < 5) {
			cont
		} else (i > 8) {
			stop
		}

		sout "The number is " + i
	}
`
	parser_code_cast = `
	sout cast::[Int](2.0)
`
	parser_code_class = `
	class Bound(constant min: Int, constant max: Int)

	class Range(constant dat: Index) {
		variable _index = 0

		function hasPrev: Bit {
			=> index != 0
		}

		function hasNext: Bit {
			=> index < dat.size
		}

		function prev: All {
			=> dat.get(index--)
		}

		function next: All {
			=> dat.get(index++)
		}
	}
`
	parser_code_trait = `
	trait Sized(constant size: Int)

	trait Index::[Sized] {
		function get(index: Int): All

		function set(index: Int, value: All)
	}
`
)

func TestParser_Parse(t *testing.T) {
	l := &Lexer{code: parser_code_property}

	l.RunLexer()

	p := &Parser{
		data: l.data,
	}

	p.RunParser()

	if err := p.fail; err != nil {
		log.Printf("Error: %+v", err)
	}

	for _, cmd := range p.cmds {
		if _, ok := cmd.(*runtime.CommandBoundary); ok {
			log.Println("=")
		} else {
			log.Println(cmd.String())
		}
	}

	// assert.NoError(t, parser.fail)
}
