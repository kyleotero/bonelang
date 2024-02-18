package lexer

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

type Token int

const (
	Tok_eof    = -1
	Tok_func   = -2
	Tok_extern = -3
	Tok_id     = -4
	Tok_number = -5
)

type Lexer struct {
	reader    *bufio.Reader
	last_char rune
}

func NewLexer() *Lexer {
	return &Lexer{
		reader:    bufio.NewReader(os.Stdin),
		last_char: ' ',
	}
}

func (l *Lexer) getChar() rune {
	input, _, err := l.reader.ReadRune()

	if err != nil {
		if err == io.EOF {
			return -1
		}

		fmt.Println("error getting char")
	}

	return input
}

func (l *Lexer) GetToken() int {

	for unicode.IsSpace(l.last_char) {
		l.last_char = l.getChar()
	}

	if unicode.IsLetter(l.last_char) {
		token_string := ""

		for unicode.IsLetter(l.last_char) || unicode.IsDigit(l.last_char) {
			token_string += string(l.last_char)
			l.last_char = l.getChar()
		}

		if token_string == "func" {
			return Tok_func
		} else if token_string == "extern" {
			return Tok_extern
		}
		return Tok_id
	}

	if unicode.IsDigit(l.last_char) || l.last_char == '.' {
		num_string := ""
		period_seen := false
		valid := true

		for unicode.IsDigit(l.last_char) || l.last_char == '.' {
			if l.last_char == '.' {
				if period_seen {
					valid = false
				} else {
					period_seen = true
				}
			}

			num_string += string(l.last_char)
			l.last_char = l.getChar()
		}

		if !valid {
			log.Fatalf("%s is not a valid number.", num_string)
		}

		return Tok_number
	}

	if l.last_char == '$' {
		for l.last_char != '\n' && l.last_char != '\r' && l.last_char != -1 {
			l.last_char = l.getChar()
		}

		if l.last_char != -1 {
			return l.GetToken()
		}
	}

	if l.last_char == -1 {
		return Tok_eof
	}

	return int(l.last_char)
}
