package words

import (
	"strings"
	"unicode"

	"github.com/kljensen/snowball"
)

type Stammer struct{}

func NewStammer() *Stammer {
	return &Stammer{}
}

func removePunctuation(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if !unicode.IsPunct(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func removeStopWords(words []string, stopWords map[string]bool) []string {
	var result []string
	for _, word := range words {
		if !stopWords[strings.ToLower(word)] {
			result = append(result, word)
		}
	}
	return result
}

func (s *Stammer) Stem(words string) ([]string, error) {
	words = removePunctuation(words)
	wordsSlice := strings.Fields(words)
	wordsSlice = removeStopWords(wordsSlice, Exclusions)

	var stammedSlice []string

	for _, word := range wordsSlice {
		stemmed, err := snowball.Stem(word, "english", true)
		if err != nil {
			continue
		}
		stammedSlice = append(stammedSlice, stemmed)
	}

	return stammedSlice, nil
}
