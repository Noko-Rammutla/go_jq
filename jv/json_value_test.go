package jv_test

import (
	"fmt"
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
		{input: "5", value: jv.NewNumber(5, "5")},
		{input: "  null   ", value: jv.NewNull()},
		{input: "{ \"name\": \"Alice\" }", value: jv.NewObject(map[string]jv.JsonValue{
			"name": jv.NewString("Alice"),
		})},
		{input: "{\"age\": 5\n }", value: jv.NewObject(map[string]jv.JsonValue{
			"age": jv.NewNumber(5.0, "5"),
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
		{input: "{ \"name\": \"Alice\",  \"transactions\": [ 1, -4.38, {\"name\": \"Bob\"} ] }", value: jv.NewObject(map[string]jv.JsonValue{
			"name": jv.NewString("Alice"),
			"transactions": jv.NewArray([]jv.JsonValue{
				jv.NewNumber(1, "1"),
				jv.NewNumber(-4.38, ""),
				jv.NewObject(map[string]jv.JsonValue{
					"name": jv.NewString("Bob"),
				}),
			}),
		})},
	}

	for n, test := range tests {
		t.Run(fmt.Sprintf("Test %d", n), func(t *testing.T) {
			value := jv.Parse(test.input)
			if !jv.Equals(test.value, value) {
				t.Errorf("Expected %s to be parsed to %v", test.input, test.value)
			}
		})
	}
}
