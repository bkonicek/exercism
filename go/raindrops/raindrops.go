package raindrops

import "strconv"

// Convert takes an integer and returns 'Pling' if it has 3 as a factor,
// 'Plang' if it has 5 as a factor, and 'Plong' if it has 7 as a factor.
// If more than 1 of the above factors apply, their respective strings are all returned.
// If neither 3, 5, or 7 are factors, the number itself is returned as a string.
func Convert(num int) string {
	if num < 3 {
		return strconv.Itoa(num)
	}
	var output string
	if num%3 == 0 {
		output += "Pling"
	}
	if num%5 == 0 {
		output += "Plang"
	}
	if num%7 == 0 {
		output += "Plong"
	}
	if output == "" {
		return strconv.Itoa(num)
	}
	return output
}
