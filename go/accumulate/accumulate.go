package accumulate

// Accumulate applies a converter function (e.g. strings.ToUpper) to a given list
// of strings, and returns the converted list
func Accumulate(str []string, c func(string) string) []string {
	for i, v := range str {
		str[i] = c(v)
	}
	return str
}
