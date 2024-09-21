package main

import (
	"github.com/Noko-Rammutla/go_jq/jv"
)

func main() {
	jv.PrettryPrint(jv.Parse("{\"name\": 5, \"children\": [null, false, true, {\"surname\": \"Bob\"}]}"))
}
