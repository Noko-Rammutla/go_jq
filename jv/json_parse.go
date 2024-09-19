package jv

import "strings"

func Parse(input string) JsonValue {
	n := consumeWhiteSpace(input)

	consumed, value := consume(input[n:])
	if consumed == -1 {
		return NewInvalid()
	}
	n += consumed

	n += consumeWhiteSpace(input[n:])
	if n != len(input) {
		return NewInvalid()
	}
	return value
}

func consume(input string) (int, JsonValue) {
	if len(input) == 0 {
		return -1, NewInvalid()
	}

	switch input[0] {
	case 'n':
		consumed := consumeLiteral(input, "null")
		if consumed == -1 {
			return -1, NewInvalid()
		} else {
			return consumed, NewNull()
		}
	case '"':
		return consumeString(input)
	case '{':
		return consumeObject(input)
	default:
		return -1, NewInvalid()
	}
}

func isWhiteSpace(char byte) bool {
	return char == ' '
}

func consumeString(input string) (int, JsonValue) {
	if len(input) < 2 {
		return -1, NewInvalid()
	}

	endIndex := strings.Index(input[1:], "\"")
	if input[0] != '"' || endIndex == -1 {
		return -1, NewInvalid()
	}

	return endIndex + 2, NewString(input[1 : endIndex+1])
}

func consumeObject(input string) (int, JsonValue) {
	if len(input) < 2 || input[0] != '{' {
		return -1, NewInvalid()
	}

	values := make(map[string]JsonValue)
	var consumed int
	var key JsonValue
	var value JsonValue

	n := 1
	for {
		if n >= len(input) {
			break
		}
		n += consumeWhiteSpace(input[n:])

		current := string(input[n])
		if current == "}" {
			n += 1
			break
		}

		consumed, key = consumeString(input[n:])
		if consumed == -1 {
			return -1, NewInvalid()
		}
		n += consumed
		n += consumeWhiteSpace(input[n:])

		consumed = consumeLiteral(input[n:], ":")
		if consumed == -1 {
			return -1, NewInvalid()
		}
		n += consumed
		n += consumeWhiteSpace(input[n:])

		consumed, value = consume(input[n:])
		if consumed == -1 {
			return -1, NewInvalid()
		}
		n += consumed
		name := key.stringValue
		if IsValid(values[name]) {
			return -1, NewInvalid()
		}

		values[name] = value
	}

	return n, NewObject(values)
}

func consumeWhiteSpace(input string) int {
	for n := 0; n < len(input); n++ {
		if isWhiteSpace(input[n]) {
			continue
		} else {
			return n
		}
	}
	return len(input)
}

func consumeLiteral(input, literal string) int {
	if !strings.HasPrefix(input, literal) {
		return -1
	}
	return len(literal)
}
