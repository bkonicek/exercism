package strand

// ToRNA takes a DNA strand and returns its RNA complement
func ToRNA(dna string) string {
	if len(dna) == 0 {
		return ""
	}
	output := ""
	for _, v := range dna {
		letter := ""
		if v == 'G' {
			letter = "C"
		}
		if v == 'C' {
			letter = "G"
		}
		if v == 'T' {
			letter = "A"
		}
		if v == 'A' {
			letter = "U"
		}
		output = output + letter
	}
	return output
}
