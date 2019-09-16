package romannumerals

import (
	"errors"
)

var conversions = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

// ToRomanNumeral takes an arabic int and convert it into a string of roman numerals
// representing the same number. Return an error if the arabic number
// is 0 or less, or greater than 3000.
func ToRomanNumeral(arabic int) (string, error) {
	if arabic <= 0 || arabic > 3000 {
		return "", errors.New("Invalid arabic numeral")
	}
	var numerals = ""
	numbers := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
beginning:
	for arabic > 0 {
		for i := 0; i < len(numbers); i++ {
			if numbers[i] <= arabic {
				arabicVal := numbers[i]
				numerals += conversions[arabicVal]
				arabic -= arabicVal
				goto beginning
			}
		}
	}
	return numerals, nil
}
