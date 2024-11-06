package eval

import "github.com/Noko-Rammutla/go_jq/jv"

func Run(input jv.JsonValue, program string) (jv.JsonValue, error) {
	tokens, err := Scan(program)
	if err != nil {
		return jv.NewInvalid(), err
	}

	filters, err := Parse(tokens)
	if err != nil {
		return jv.NewInvalid(), err
	}

	var output = input
	for _, filter := range filters {
		output, err = filter.Apply(output)
		if err != nil {
			return jv.NewInvalid(), err
		}
	}

	return output, nil
}
