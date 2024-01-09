package scrabble

// Source: exercism/problem-specifications
// Commit: 0d882ed scrabble-score: Apply new "input" policy
// Problem Specifications Version: 1.1.0

type scrabbleTest struct {
	input       string
	expected    int
	description string
}

var scrabbleScoreTests = []scrabbleTest{
	{"a", 1, "foo"},                           // lowercase letter
	{"A", 1, "foo"},                           // uppercase letter
	{"f", 4, "foo"},                           // valuable letter
	{"at", 2, "foo"},                          // short word
	{"zoo", 12, "foo"},                        // short, valuable word
	{"street", 6, "foo"},                      // medium word
	{"quirky", 22, "foo"},                     // medium, valuable word
	{"OxyphenButazone", 41, "foo"},            // long, mixed-case word
	{"pinata", 8, "foo"},                      // english-like word
	{"", 0, "foo"},                            // empty input
	{"abcdefghijklmnopqrstuvwxyz", 87, "foo"}, // entire alphabet available
}
