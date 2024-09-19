package jv_test

import (
	"testing"

	"github.com/Noko-Rammutla/go_jq/jv"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input string
		value jv.JsonValue
	}{
		{input: "", value: jv.NewInvalid()},
		{input: "   ", value: jv.NewInvalid()},
		{input: "\"Alice\"", value: jv.NewString("Alice")},
		{input: " \"Bob\"   ", value: jv.NewString("Bob")},
		{input: "null", value: jv.NewNull()},
		{input: "  null   ", value: jv.NewNull()},
		{input: "{ \"name\": \"Alice\" }", value: jv.NewObject(map[string]jv.JsonValue{
			"name": jv.NewString("Alice"),
		})},
		{input: "{ \"true\": true,  \"false\": false }", value: jv.NewObject(map[string]jv.JsonValue{
			"true":  jv.NewBoolean(true),
			"false": jv.NewBoolean(false),
		})},
		{input: "{ \"name\": \"Alice\",  \"child\": { \"name\": \"Bob\" } }", value: jv.NewObject(map[string]jv.JsonValue{
			"name": jv.NewString("Alice"),
			"child": jv.NewObject(map[string]jv.JsonValue{
				"name": jv.NewString("Bob"),
			}),
		})},
	}

	for _, test := range tests {
		value := jv.Parse(test.input)
		if !jv.Equals(test.value, value) {
			t.Errorf("Expected %s to be parsed to %v", test.input, test.value)
		}
	}
}
