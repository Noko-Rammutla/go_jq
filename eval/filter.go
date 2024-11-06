package eval

import (
	"fmt"
	"github.com/Noko-Rammutla/go_jq/jv"
)

type Filter interface {
	Apply(input jv.JsonValue) (jv.JsonValue, error)
}

type Identity struct{}

func NewIdentity() Filter {
	return Identity{}
}

func (op Identity) Apply(input jv.JsonValue) (jv.JsonValue, error) {
	return input, nil
}

type ObjectIndex struct {
	identifier string
}

func NewObjectIndex(identifier string) Filter {
	return ObjectIndex{identifier: identifier}
}

func (op ObjectIndex) Apply(input jv.JsonValue) (jv.JsonValue, error) {
	if input.GetKind() != jv.JObject {
		return jv.NewInvalid(), fmt.Errorf("expected object but found %s", input.GetKind())
	}
	object := input.GetAsObject()
	if value, exists := object[op.identifier]; exists {
		return value, nil
	}
	return jv.NewInvalid(), fmt.Errorf("%s does not exist in object", op.identifier)
}
