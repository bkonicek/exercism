package scrabble

import (
	"strings"
)

func Score(word string) int {
	word = strings.ToUpper(word)
	var total int
	// For each character, find its value and multiply by
	// The number of times it occurs in the string
	for _, w := range word {
		switch w {
		case 'D', 'G':
			total += 2
		case 'B', 'C', 'M', 'P':
			total += 3
		case 'F', 'H', 'V', 'W', 'Y':
			total += 4
		case 'K':
			total += 5
		case 'J', 'X':
			total += 8
		case 'Q', 'Z':
			total += 10
		default:
			total += 1
		}
	}
	return total
}
