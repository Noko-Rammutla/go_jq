package jv

import (
	"strings"
)

type JsonValue struct {
	stringValue string
	objectValue map[string]JsonValue
	kind        JsonKind
}

type JsonKind string

const (
	JString  JsonKind = "string"
	JNull    JsonKind = "null"
	JInvalid JsonKind = "invalid"
	JObject  JsonKind = "object"
)

func NewString(value string) JsonValue {
	return JsonValue{
		stringValue: value,
		kind:        JString,
	}
}

func NewObject(value map[string]JsonValue) JsonValue {
	return JsonValue{
		objectValue: value,
		kind:        JObject,
	}
}

func NewNull() JsonValue {
	return JsonValue{
		kind: JNull,
	}
}

func NewInvalid() JsonValue {
	return JsonValue{
		kind: JInvalid,
	}
}

func Equals(lhs, rhs JsonValue) bool {
	switch lhs.kind {
	case JInvalid:
		return rhs.kind == JInvalid
	case JNull:
		return rhs.kind == JNull
	case JString:
		return rhs.kind == JString && lhs.stringValue == rhs.stringValue
	case JObject:
		return rhs.kind == JObject && objectEquals(lhs, rhs)
	}
	return false
}

func IsValid(value JsonValue) bool {
	if value.kind == "" {
		return false
	}
	return value.kind != JInvalid
}

func objectEquals(lhs, rhs JsonValue) bool {
	if len(lhs.objectValue) != len(rhs.objectValue) {
		return false
	}

	for k := range lhs.objectValue {
		if !Equals(lhs.objectValue[k], rhs.objectValue[k]) {
			return false
		}
	}

	return true
}

func Parse(input string) JsonValue {
	n := 0
	value := NewInvalid()
	consumed := -1

	for {
		if n >= len(input) {
			break
		}
		if isWhiteSpace(input[n]) {
			n += 1
			continue
		} else if IsValid(value) {
			return NewInvalid()
		} else {
			consumed, value = parse(input[n:])
			if consumed == -1 {
				return NewInvalid()
			}
			n += consumed
			continue
		}
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

		current := string(input[n])
		if current == "}" {
			n += 1
			break
		}

		if isWhiteSpace(current[0]) {
			n += 1
			continue
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
