package lexer

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

type token int

const (
	tok_eof    = -1
	tok_func   = -2
	tok_extern = -3
	tok_id     = -4
	tok_number = -5
)

func get_char(reader bufio.Reader) rune {
	input, _, err := reader.ReadRune()

	if err != nil {
		if err == io.EOF {
			return -1
		}

		fmt.Println("error getting char")
	}

	return input
}

func get_tok() int {
	last_char := ' '

	reader := bufio.NewReader(os.Stdin)

	for unicode.IsSpace(last_char) {
		last_char = get_char(*reader)
	}

	if unicode.IsLetter(last_char) {
		token_string := ""

		for unicode.IsLetter(last_char) && unicode.IsDigit(last_char) {
			token_string += string(last_char)
			last_char = get_char(*reader)
		}

		if token_string == "func" {
			return tok_func
		} else if token_string == "extern" {
			return tok_extern
		}
		return tok_id
	}

	if unicode.IsDigit(last_char) || last_char == '.' {
		num_string := ""
		period_seen := false
		valid := true

		for unicode.IsDigit(last_char) || last_char == '.' {
			if last_char == '.' {
				if period_seen {
					valid = false
				} else {
					period_seen = true
				}
			}

			num_string += string(last_char)
			last_char = get_char(*reader)
		}

		if !valid {
			log.Fatalf("%s is not a valid number.", num_string)
		}

		return tok_number
	}

	if last_char == '$' {
		for last_char != '\n' && last_char != '\r' && last_char != -1 {
			last_char = get_char(*reader)
		}

		if last_char != -1 {
			return get_tok()
		}
	}

	if last_char == -1 {
		return tok_eof
	}

	return int(last_char)
}
