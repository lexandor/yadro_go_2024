package main

import "strings"

var PrefixToReplace = map[string]string{
	"un":  "",
	"dis": "",
	"in":  "",
	"im":  "",
	"ir":  "",
	"ll":  "",
	"pre": "",
	"re":  "",
}

func RemovePrefix(word, suffix, replace string) string {
	if strings.HasPrefix(word, suffix) {
		return strings.TrimPrefix(word, suffix) + replace
	}

	return word
}
