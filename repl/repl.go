package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ellemouton/monkey/parser"

	"github.com/ellemouton/monkey/lexer"
)

/* REPL: Read Eval Print Loop

Sometimes called "console" or "interactive mode".
The REPL reads input, sends it to the interpreter for evaluation,
prints the result/output of the interpreter and starts again.

*/

const PROMPT = ">> "

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

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
