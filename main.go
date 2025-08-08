package main

import (
	in "BFG/interpreter"
	lex "BFG/lexer"
	lc "BFG/loopcount"
	// "fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		panic("No source code given")
	}
	code, err := os.ReadFile(os.Args[1])

	if err != nil {
		panic(err)
	}

	loopcount := lc.New()
	lexer := lex.New(code)
	interpreter := in.New(&loopcount, &lexer)

	token := lexer.GetToken()

	for token.Kind != lex.EOF && lexer.CurPos < len(lexer.Source) {
		interpreter.HandleToken(token)
		token = lexer.GetToken()
	}
}
