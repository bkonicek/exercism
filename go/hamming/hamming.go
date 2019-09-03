package hamming

import "errors"

// Distance counts the number of differences at a single
// nucleotide between two DNA strands.
//
// If they are not equal lengths, it will return true for the error value.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings are not equal length")
	}
	var diffs int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffs++
		}
	}
	return diffs, nil
}
