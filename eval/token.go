package eval

type TokenType string

const (
	DOT          TokenType = "."
	PIPE         TokenType = "|"
	LEFT_SQUARE  TokenType = "["
	RIGHT_SQUARE TokenType = "]"
	INTEGER      TokenType = "integer"
	IDENTIFIER   TokenType = "identifier"
	INVALID      TokenType = "invalid"
)

type Token struct {
	tokenType TokenType
	value     interface{}
}

func NewToken(tokenType TokenType) Token {
	return Token{
		tokenType: tokenType,
	}
}

func NewIntegerToken(value int) Token {
	return Token{
		tokenType: INTEGER,
		value:     value,
	}
}

func NewIdentifier(name string) Token {
	return Token{
		tokenType: IDENTIFIER,
		value:     name,
	}
}
