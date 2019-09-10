package raindrops

import "strconv"

const three = "Pling"
const five = "Plang"
const seven = "Plong"

// Convert takes an integer and returns 'Pling' if it has 3 as a factor,
// 'Plang' if it has 5 as a factor, and 'Plong' if it has 7 as a factor.
// If more than 1 of the above factors apply, their respective strings are all returned.
// If neither 3, 5, or 7 are factors, the number itself is returned as a string.
func Convert(num int) string {
	if num < 3 {
		return strconv.Itoa(num)
	}
	max := num
	if num > 7 {
		max = 7
	}
	var output string
	for i := 3; i <= max; i = i + 2 {
		if num%i == 0 {
			if i == 3 {
				output += three
			}
			if i == 5 {
				output += five
			}
			if i == 7 {
				output += seven
			}
		}
	}
	if output == "" {
		return strconv.Itoa(num)
	}
	return output
}
