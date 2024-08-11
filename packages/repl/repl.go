package repl

import (
	"bufio"
	"fmt"
	"github.com/limistah/monlang/packages/evaluator"
	"github.com/limistah/monlang/packages/parser"
	"io"

	"github.com/limistah/monlang/packages/lexer"
)

const PROMPT = ">>>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printIOError(err error) {
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func printParserErrors(out io.Writer, errors []string) {
	_, err := io.WriteString(out, "Whoops! We ran into some monkey business here!\n")
	if err != nil {
		printIOError(err)
	}
	for _, msg := range errors {
		_, err = io.WriteString(out, "\t"+msg+"\n")
		if err != nil {
			printIOError(err)
		}
	}
	_, err = io.WriteString(out, "\n")
	if err != nil {
		printIOError(err)
	}
}
