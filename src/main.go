package main

import (
	"fmt"
	"llvm-lang/src/lexer"
)

func main() {
	// Initialize lexer with os.Stdin
	// Loop until EOF token is encountered

	l := lexer.NewLexer()

	for tok := l.GetToken(); tok != lexer.Tok_eof; tok = l.GetToken() {
		fmt.Printf("Token: %v\n", tok)
	}
}
