// Package acronym implements a routine to convert a phrase to its acronym.
package acronym

import (
	"regexp"
	"strings"
)

// Regular expression to find individual words in the string.
var re = regexp.MustCompile(`([[:alpha:]]+('s)?)`)

// Abbreviate accepts a string input s and returns the acronym for it.
func Abbreviate(s string) string {
	if len(s) == 0 {
		return ""
	}
	words := re.FindAllString(s, -1)
	var acronym string
	for _, v := range words {
		acronym += strings.ToUpper(v[0:1])
	}
	return acronym
}
