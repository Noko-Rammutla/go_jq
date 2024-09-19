package jv

import "strings"

type JsonValue struct {
	stringValue string
	kind        JsonKind
}

type JsonKind string

const (
	JString  JsonKind = "string"
	JNull    JsonKind = "null"
	JInvalid JsonKind = "invalid"
)

func NewString(value string) JsonValue {
	return JsonValue{
		stringValue: value,
		kind:        JString,
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
	}
	return false
}

func Parse(value string) JsonValue {
	n := 0
	var current byte
	for {
		if n >= len(value) {
			break
		}
		current = value[n]

		if isWhiteSpace(current) {
			n += 1
			continue
		}

		switch current {
		case 'n':
			return parseNull(value[n:])
		case '"':
			return parseString(value[n:])
		default:
			return NewInvalid()
		}
	}

	return NewInvalid()
}

func isWhiteSpace(char byte) bool {
	return char == ' '
}

func parseNull(value string) JsonValue {
	if !strings.HasPrefix(value, "null") {
		return NewInvalid()
	}
	for n := 4; n < len(value); n++ {
		if !isWhiteSpace(value[n]) {
			return NewInvalid()
		}
	}
	return NewNull()
}

func parseString(value string) JsonValue {
	foundEnd := false
	endIndex := 0

	for n := 1; n < len(value); n++ {
		if foundEnd && !isWhiteSpace(value[n]) {
			return NewInvalid()
		}
		if value[n] == '"' {
			foundEnd = true
			endIndex = n
			continue
		}
	}

	if !foundEnd {
		return NewInvalid()
	}
	return NewString(value[1:endIndex])
}
