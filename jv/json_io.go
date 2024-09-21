package jv

import (
	"fmt"
	"strconv"
	"strings"
)

func PrettryPrint(value JsonValue) {
	fmt.Println(formatJson(value, 0))
}

func formatJson(value JsonValue, offset int) string {
	switch value.kind {
	case JArray:
		return formatArray(value, offset)
	case JObject:
		return formatObject(value, offset)
	default:
		return formatSimple(value)
	}
}

func formatArray(value JsonValue, offset int) string {
	prefix := strings.Repeat("  ", offset)
	text := "["

	for n := range value.arrayValue {
		text += fmt.Sprintf("\n  %s%s,", prefix, formatJson(value.arrayValue[n], offset+1))
	}

	if text[len(text)-1] == ',' {
		text = text[:len(text)-1]
	}

	text += fmt.Sprintf("\n%s]", prefix)
	return text
}

func formatObject(value JsonValue, offset int) string {
	prefix := strings.Repeat("  ", offset)
	text := "{"

	for key := range value.objectValue {
		valueText := formatJson(value.objectValue[key], offset+1)
		keyText := colourText(fmt.Sprintf("\"%s\"", key), ColourCyan)
		text += fmt.Sprintf("\n%s  %s : %s,", prefix, keyText, valueText)
	}

	if text[len(text)-1] == ',' {
		text = text[:len(text)-1]
	}

	text += fmt.Sprintf("\n%s}", prefix)
	return text
}

func formatSimple(value JsonValue) string {
	switch value.kind {
	case JTrue:
		return colourText(string(value.kind), ColourBlue)
	case JNumber:
		numberText := value.stringValue
		if numberText == "" {
			numberText = strconv.FormatFloat(value.numberValue, 'f', 4, 64)
		}
		return colourText(numberText, ColourMagenta)
	case JFalse, JNull:
		return colourText(string(value.kind), ColourRed)
	case JString:
		return colourText(fmt.Sprintf("\"%s\"", value.stringValue), ColourGreen)
	}

	return ""
}

type ColourCode int8

const (
	ColourRed     ColourCode = 91
	ColourGreen   ColourCode = 92
	ColourYellow  ColourCode = 93
	ColourBlue    ColourCode = 94
	ColourMagenta ColourCode = 95
	ColourCyan    ColourCode = 96
)

func colourText(text string, colourCode ColourCode) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", colourCode, text)
}
