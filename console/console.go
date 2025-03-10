package console

/* Provides the standard REPL+ for Mika */

import (
	"runtime"
	"bufio"
	"io"
	"fmt"
	"github.com/aenlemmea/mika/front/lexer"
	"github.com/aenlemmea/mika/front/token"
)

const PROMPT = "#>> "

func Console(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	
	for {
		fmt.Printf(PROMPT)
		read := scanner.Scan()
		if !read {
			return
		}
		text := scanner.Text()
		if text == "exit" {
			return
		} else if text == "clear" {
			if (runtime.GOOS == "linux") {
				fmt.Printf("\033[H\033[2J")
				continue
			}
		}
		lex := lexer.New(text)
		
		// TODO Avoid relying on the first token for loop init
		for tok := lex.NextToken(); tok.Kind != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
