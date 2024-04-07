package main

import "strings"

var SuffixesToReplace = map[string]string{
	"ies": "y",
	"es":  "",
	"s":   "",
	"er":  "",
	"ing": "",
	"ed":  "",
	",":   "",
	".":   "",
	"!":   "",
	"?":   "",
	":":   "",
	";":   "",
}

func RemoveSuffix(word, suffix, replace string) string {
	if strings.HasSuffix(word, suffix) {
		return strings.TrimSuffix(word, suffix) + replace
	}

	return word
}
