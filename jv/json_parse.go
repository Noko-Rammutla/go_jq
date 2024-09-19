package jv

import "strings"

func Parse(input string) JsonValue {
	n := consumeWhiteSpace(input)

	consumed, value := parse(input[n:])
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

func parse(input string) (int, JsonValue) {
	if len(input) == 0 {
		return -1, NewInvalid()
	}

	switch input[0] {
	case 'n':
		return parseNull(input)
	case '"':
		return parseString(input)
	case '{':
		return parseObject(input)
	default:
		return -1, NewInvalid()
	}
}

func isWhiteSpace(char byte) bool {
	return char == ' '
}

func parseNull(input string) (int, JsonValue) {
	if !strings.HasPrefix(input, "null") {
		return -1, NewInvalid()
	}
	return 4, NewNull()
}

func parseString(input string) (int, JsonValue) {
	if len(input) < 2 {
		return -1, NewInvalid()
	}

	endIndex := strings.Index(input[1:], "\"")
	if input[0] != '"' || endIndex == -1 {
		return -1, NewInvalid()
	}

	return endIndex + 2, NewString(input[1 : endIndex+1])
}

func parseObject(input string) (int, JsonValue) {
	if len(input) < 2 || input[0] != '{' {
		return -1, NewInvalid()
	}

	values := make(map[string]JsonValue)
	var consumed int
	var key JsonValue
	var value JsonValue
	foundSeperator := false

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

		if !IsValid(key) {
			consumed, key = parseString(input[n:])
			if consumed == -1 {
				return -1, NewInvalid()
			}
			n += consumed
			continue
		} else if !foundSeperator && current == ":" {
			foundSeperator = true
			n += 1
			continue
		} else if foundSeperator {
			consumed, value = parse(input[n:])
			if consumed == -1 {
				return -1, NewInvalid()
			}

			name := key.stringValue
			if IsValid(values[name]) {
				return -1, NewInvalid()
			}

			values[name] = value
			key = NewInvalid()
			value = NewInvalid()
			n += consumed
			continue
		} else {
			return -1, NewInvalid()
		}
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
