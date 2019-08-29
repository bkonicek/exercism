package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	var n string
	if name == "" {
		n = "you"
	} else {
		n = name
	}
	return fmt.Sprintf("One for %v, one for me.", n)
}
