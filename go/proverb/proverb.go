package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	var output []string

	if len(rhyme) != 0 {
		for index := 0; index < len(rhyme)-1; index++ {
			output = append(output, fmt.Sprintf("For want of a %v the %v was lost.", rhyme[index], rhyme[index+1]))
		}
		output = append(output, fmt.Sprintf("And all for the want of a %v.", rhyme[0]))
	}
	return output
}
