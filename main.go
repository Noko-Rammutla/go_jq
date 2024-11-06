package main

import (
	_ "embed"
	"fmt"
	"github.com/Noko-Rammutla/go_jq/eval"
	"github.com/Noko-Rammutla/go_jq/jv"
	"io"
	"os"
)

//go:embed data/help.txt
var help string

func main() {
	args := os.Args[1:]

	if len(args) > 2 {
		_, err := fmt.Fprint(os.Stderr, help)
		exitOnError(err)
		os.Exit(1)
	}

	var source []byte
	var err error
	if len(args) == 2 {
		_, err = os.Stat(args[1])
		if os.IsNotExist(err) {
			_, err = fmt.Fprintf(os.Stderr, "File not found '%s'\n", args[1])
			exitOnError(err)
			os.Exit(1)
		}
		source, err = os.ReadFile(args[1])
		if err != nil {
			_, err = fmt.Fprintf(os.Stderr, "Failed to read file '%s'\n", args[1])
			exitOnError(err)
			os.Exit(1)
		}
	} else {
		source, err = io.ReadAll(os.Stdin)
		exitOnError(err)
	}

	inputValue := jv.Parse(string(source))

	var filter string
	if len(args) > 0 {
		filter = args[0]
	} else {
		filter = "."
	}

	outputValue, err := eval.Run(inputValue, filter)
	exitOnError(err)

	jv.PrettyPrint(outputValue)
}

func exitOnError(err error) {
	if err == nil {
		return
	}
	_, writeErr := fmt.Fprintf(os.Stderr, err.Error())
	if writeErr != nil {
		panic(err)
	}
	os.Exit(1)
}
