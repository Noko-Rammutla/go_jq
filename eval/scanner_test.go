package eval_test

import (
	"github.com/Noko-Rammutla/go_jq/eval"
	"testing"
)

func assertEquals(t *testing.T, expected, actual []eval.Token) {
	if len(expected) != len(actual) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
	}
	for n := range expected {
		if expected[n] != actual[n] {
			t.Fatalf("Expected '%v', got '%v' at position %d", expected[n], actual[n], n)
		}
	}
}

func TestScan(t *testing.T) {
	tests := map[string]struct {
		input  string
		output []eval.Token
	}{
		"dot": {
			input:  ".",
			output: []eval.Token{eval.NewToken(eval.DOT)},
		},
		"pipe": {
			input:  "|",
			output: []eval.Token{eval.NewToken(eval.PIPE)},
		},
		"array index": {
			input:  "[-17]",
			output: []eval.Token{eval.NewToken(eval.LEFT_SQUARE), eval.NewIntegerToken(-17), eval.NewToken(eval.RIGHT_SQUARE)},
		},
		"object index": {
			input:  ".name",
			output: []eval.Token{eval.NewToken(eval.DOT), eval.NewIdentifier("name")},
		},
		"mixed": {
			input: ".name| [23] |.",
			output: []eval.Token{
				eval.NewToken(eval.DOT),
				eval.NewIdentifier("name"),
				eval.NewToken(eval.PIPE),
				eval.NewToken(eval.LEFT_SQUARE),
				eval.NewIntegerToken(23),
				eval.NewToken(eval.RIGHT_SQUARE),
				eval.NewToken(eval.PIPE),
				eval.NewToken(eval.DOT),
			},
		},
	}

	for name, data := range tests {
		t.Run(name, func(t *testing.T) {
			output, err := eval.Scan(data.input)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
				return
			}
			assertEquals(t, data.output, output)
		})
	}
}
