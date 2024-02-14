package parser

import (
	"fmt"
	"strings"
)

var traceLevel int = 0

const traceIndentPlaceholder = "\t"

func identtraceLevelevel() string {
	return strings.Repeat(traceIndentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identtraceLevelevel(), fs)
}

func incIdent() { traceLevel = traceLevel + 1 }
func decIdent() { traceLevel = traceLevel - 1 }

func trace(msg string) string {
	incIdent()
	tracePrint("BEGIN: " + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END: " + msg)
	decIdent()
}
