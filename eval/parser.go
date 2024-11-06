package eval

import "fmt"

type parser struct {
	tokens  []Token
	current int
}

func Parse(tokens []Token) ([]Filter, error) {
	if len(tokens) == 0 {
		return []Filter{
			NewIdentity(),
		}, nil
	}
	p := parser{
		tokens:  tokens,
		current: 0,
	}
	return p.parse()
}

func (p *parser) parse() ([]Filter, error) {
	filters := make([]Filter, 0, len(p.tokens))
	for {
		filter, err := p.expression()
		if err != nil {
			return nil, err
		}
		filters = append(filters, filter)

		if p.isAtEnd() {
			break
		}

		err = p.consume(PIPE)
		if err != nil {
			return nil, err
		}
	}

	return filters, nil
}

func (p *parser) expression() (Filter, error) {
	token := p.advance()
	switch token.tokenType {
	case DOT:
		if identifier, found := p.match(IDENTIFIER); found {
			return NewObjectIndex(identifier.value.(string)), nil
		}
		return NewIdentity(), nil
	default:
		return nil, fmt.Errorf("failed to parse input at %s", token.tokenType)
	}
}

func (p *parser) isAtEnd() bool {
	return p.current >= len(p.tokens)
}

func (p *parser) peek() Token {
	if p.isAtEnd() {
		return NewToken(INVALID)
	}
	return p.tokens[p.current]
}

func (p *parser) advance() Token {
	if p.isAtEnd() {
		return NewToken(INVALID)
	}
	p.current += 1
	return p.tokens[p.current-1]
}

func (p *parser) match(tokentype TokenType) (Token, bool) {
	if p.peek().tokenType == tokentype {
		return p.advance(), true
	}
	return NewToken(INVALID), false
}

func (p *parser) consume(tokenType TokenType) error {
	if p.isAtEnd() || p.peek().tokenType != tokenType {
		return fmt.Errorf("expected to find %s at %d", tokenType, p.current)
	}
	p.current += 1
	return nil
}
