package eval_test

import (
	"github.com/Noko-Rammutla/go_jq/eval"
	"testing"
)

func tokenEqual(lhs, rhs []eval.Token) bool {
	return false
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
			input:  "[0]",
			output: []eval.Token{eval.NewToken(eval.LEFT_SQUARE), eval.NewIntegerToken(0), eval.NewToken(eval.RIGHT_SQUARE)},
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
			if !tokenEqual(output, data.output) {
				t.Fatal("Unexpected output")
			}
		})
	}
}
