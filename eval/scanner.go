package eval

import (
	"fmt"
	"strconv"
)

func Scan(source string) ([]Token, error) {
	s := scanner{
		source:  source,
		current: 0,
	}
	return s.scan()
}

type scanner struct {
	source  string
	current int
}

func (s *scanner) scan() ([]Token, error) {
	tokens := make([]Token, 0, len(s.source))
	for {
		if s.isAtEnd() {
			break
		}
		c := s.advance()
		switch c {
		case '.', '|', '[', ']':
			tokens = append(tokens, NewToken(TokenType(c)))
			break
		default:
			if isWhiteSpace(c) {
				break
			} else if isDigit(c) || c == '-' {
				value, err := s.consumeInteger(c)
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, NewIntegerToken(value))
			} else if isAlpha(c) {
				value := s.consumeIdentifier(c)
				tokens = append(tokens, NewIdentifier(value))
			}
		}
	}
	return tokens, nil
}

func (s *scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}

func (s *scanner) advance() byte {
	if s.isAtEnd() {
		return 0
	}
	s.current += 1
	return s.source[s.current-1]
}

func (s *scanner) error(value byte) error {
	return fmt.Errorf("invalid character '%c' at %d", value, s.current)
}

func (s *scanner) consumeInteger(first byte) (int, error) {
	value := "" + string(first)
	for {
		c := s.peek()
		if isDigit(c) {
			value += string(c)
			s.advance()
		} else {
			break
		}
	}
	if value == "-" {
		return 0, s.error(first)
	}
	return strconv.Atoi(value)
}

func (s *scanner) consumeIdentifier(first byte) string {
	value := "" + string(first)
	for {
		c := s.peek()
		if isDigit(c) || isAlpha(c) {
			value += string(c)
			s.advance()
		} else {
			break
		}
	}
	return value
}

func isDigit(value byte) bool {
	return value >= '0' && value <= '9'
}

func isAlpha(value byte) bool {
	return (value >= 'a' && value <= 'z') || (value >= 'A' && value <= 'Z')
}

func isWhiteSpace(value byte) bool {
	return value == ' '
}
