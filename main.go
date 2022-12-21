package main

import (
	"fmt"
	"log"
	"os"
	"sync/atomic"
	"time"

	"tanna/commons/token"
	"tanna/compile"
	"tanna/runtime"
)

const (
	do_async_parsing_depth_ticker = true
)

func main() {
	fmt.Println("\n\n")

	code, err := os.ReadFile("loop.tanna")
	if err != nil {
		log.Printf("File Read Error: %+v", err)
		return
	}

	cmds, err := Compile(string(code))
	if err != nil {
		log.Printf("Compilation Error: %+v", err)
		return
	}

	fmt.Println("\n\n\n")
	log.Println("==== Tanna Evaluation")

	stack, context, err := runtime.Execute(cmds)
	if err != nil {
		log.Printf("Execution Error: %+v\n", err)
	}

	log.Println("==== Tanna Evaluation Complete")
	fmt.Println("\n\n\n")

	log.Printf("Runtime Stack:\n%s\n", stack)
	log.Printf("Runtime State:\n%s\n", context)
}

func Compile(code string) (cmds []runtime.Command, err error) {
	l := compile.NewLexer(code)

	l.RunLexer()

	// PrintLexerData(l.LexerData())

	p := compile.NewParser(l.LexerData())

	if !do_async_parsing_depth_ticker {
		p.RunParser()
	} else {
		cancel := atomic.Bool{}
		cancel.Store(false)

		go func() {
			for range time.Tick(time.Second) {
				log.Printf("Parsing Depth %d", p.ParsingDepth())

				if cancel.Load() {
					break
				}
			}
		}()

		p.RunParser()
		cancel.Store(true)
	}

	if p.Err() != nil {
		return nil, p.Err()
	} else {
		PrintParserData(p.ParserData())

		return p.ParserData(), nil
	}
}

func PrintLexerData(tokens []*token.Data) {
	for _, data := range tokens {
		fmt.Println(data.String())
	}
}

func PrintParserData(cmds []runtime.Command) {
	for _, cmd := range cmds {
		if _, ok := cmd.(*runtime.CommandBoundary); ok {
			fmt.Println("=")
		} else {
			fmt.Println(cmd.String())
		}
	}
}
