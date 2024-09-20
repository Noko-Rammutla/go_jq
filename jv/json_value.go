package jv

type JsonValue struct {
	stringValue string
	numberValue float64
	objectValue map[string]JsonValue
	arrayValue  []JsonValue
	kind        JsonKind
}

type JsonKind string

const (
	JString  JsonKind = "string"
	JNull    JsonKind = "null"
	JInvalid JsonKind = "invalid"
	JObject  JsonKind = "object"
	JTrue    JsonKind = "true"
	JFalse   JsonKind = "false"
	JNumber  JsonKind = "number"
	JArray   JsonKind = "array"
)

func NewString(value string) JsonValue {
	return JsonValue{
		stringValue: value,
		kind:        JString,
	}
}

func NewNumber(value float64, serialised string) JsonValue {
	return JsonValue{
		numberValue: value,
		stringValue: serialised,
		kind:        JNumber,
	}
}

func NewObject(value map[string]JsonValue) JsonValue {
	return JsonValue{
		objectValue: value,
		kind:        JObject,
	}
}

func NewArray(value []JsonValue) JsonValue {
	return JsonValue{
		arrayValue: value,
		kind:       JArray,
	}
}

func NewNull() JsonValue {
	return JsonValue{
		kind: JNull,
	}
}

func NewBoolean(input bool) JsonValue {
	if input {
		return JsonValue{
			kind: JTrue,
		}
	} else {
		return JsonValue{
			kind: JFalse,
		}
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
	case JFalse:
		return rhs.kind == JFalse
	case JTrue:
		return rhs.kind == JTrue
	case JString:
		return rhs.kind == JString && lhs.stringValue == rhs.stringValue
	case JNumber:
		return rhs.kind == JNumber && lhs.numberValue == rhs.numberValue
	case JArray:
		return rhs.kind == JArray && arrayEquals(lhs, rhs)
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

func arrayEquals(lhs, rhs JsonValue) bool {
	if len(lhs.arrayValue) != len(rhs.arrayValue) {
		return false
	}

	for n := range lhs.arrayValue {
		if !Equals(lhs.arrayValue[n], rhs.arrayValue[n]) {
			return false
		}
	}

	return true
}
