package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	s := flag.String("s", "", "A string of words to process stamming")
	flag.Parse()

	var steams map[string]bool

	if *s != "" {
		steams = Steam(strings.Fields(*s))
	} else {
		steams = Steam(flag.Args())
	}

	for steam := range steams {
		fmt.Print(steam + " ")
	}

	fmt.Println()
}

func Steam(words []string) map[string]bool {
	steams := make(map[string]bool)

	for _, word := range words {
		word = strings.ToLower(word)
		if _, ok := Exclusions[word]; !ok {
			steams[stemWord(word)] = true
		}
	}

	return steams
}

func stemWord(word string) string {
	for prefix, replace := range PrefixToReplace {
		word = RemovePrefix(word, prefix, replace)
	}

	for suffix, replace := range SuffixesToReplace {
		word = RemoveSuffix(word, suffix, replace)
	}

	return word
}
