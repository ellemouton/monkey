package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ellemouton/monkey/lexer"
	"github.com/ellemouton/monkey/token"
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
