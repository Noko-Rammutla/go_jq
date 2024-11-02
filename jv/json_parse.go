package jv

import (
	"strconv"
	"strings"
)

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
	case 'f':
		consumed := consumeLiteral(input, "false")
		if consumed == -1 {
			return -1, NewInvalid()
		} else {
			return consumed, NewBoolean(false)
		}
	case 't':
		consumed := consumeLiteral(input, "true")
		if consumed == -1 {
			return -1, NewInvalid()
		} else {
			return consumed, NewBoolean(true)
		}
	case '"':
		return consumeString(input)
	case '{':
		return consumeObject(input)
	case '[':
		return consumeArray(input)
	default:
		return consumeNumber(input)
	}
}

func isWhiteSpace(char byte) bool {
	return char == ' ' || char == '\n' || char == '\t'
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

func consumeNumber(input string) (int, JsonValue) {
	if len(input) == 0 {
		return -1, NewInvalid()
	}
	endIndex := findNumberEnd(input)
	if endIndex == -1 {
		endIndex = len(input)
	}

	value, err := strconv.ParseFloat(input[:endIndex], 64)
	if err != nil {
		return -1, NewInvalid()
	}

	return endIndex, NewNumber(value, input[:endIndex])
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
		name := key.stringValue
		if IsValid(values[name]) {
			return -1, NewInvalid()
		}
		values[name] = value
		n += consumed
		n += consumeWhiteSpace(input[n:])

		current := input[n]
		if current == '}' {
			n += 1
			break
		} else if current == ',' {
			n += 1
			continue
		}
	}

	return n, NewObject(values)
}

func consumeArray(input string) (int, JsonValue) {
	if len(input) < 2 || input[0] != '[' {
		return -1, NewInvalid()
	}

	values := make([]JsonValue, 0)
	var consumed int
	var value JsonValue

	n := 1
	for {
		if n >= len(input) {
			break
		}
		n += consumeWhiteSpace(input[n:])

		consumed, value = consume(input[n:])
		if consumed == -1 {
			return -1, NewInvalid()
		}
		values = append(values, value)
		n += consumed
		n += consumeWhiteSpace(input[n:])

		current := input[n]
		if current == ']' {
			n += 1
			break
		} else if current == ',' {
			n += 1
			continue
		}
	}

	return n, NewArray(values)
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

func findNumberEnd(input string) int {
	for n := 0; n < len(input); n++ {
		current := string(input[n])
		if !strings.ContainsAny(current, ",{}[]\"") && !isWhiteSpace(input[n]) {
			continue
		} else {
			return n
		}
	}
	return -1
}

func consumeLiteral(input, literal string) int {
	if !strings.HasPrefix(input, literal) {
		return -1
	}
	return len(literal)
}
