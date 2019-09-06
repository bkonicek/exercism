package raindrops

import "strconv"

// FindFactors returns an array of all factors for a given integer
func FindFactors(num int) []int {
	var factors []int
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

// Convert takes an integer and returns 'Pling' if it has 3 as a factor,
// 'Plang' if it has 5 as a factor, and 'Plong' if it has 7 as a factor.
// If more than 1 of the above factors apply, their respective strings are all returned.
// If neither 3, 5, or 7 are factors, the number itself is returned as a string.
func Convert(num int) string {
	factors := FindFactors(num)
	var output string
	for i := 0; i < len(factors); i++ {
		if factors[i] == 3 {
			output = output + "Pling"
		} else if factors[i] == 5 {
			output = output + "Plang"
		} else if factors[i] == 7 {
			output = output + "Plong"
		}
	}
	if output == "" {
		return strconv.Itoa(num)
	}
	return output
}
