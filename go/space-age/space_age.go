package space

// Planet type holds the name of a planet
type Planet string

// Age converts a person'sage in seconds into
// what their age would be on a given planet
func Age(seconds float64, planet Planet) float64 {
	base := seconds / 31557600.00
	var age float64
	switch planet {
	case "Earth":
		age = base
	case "Mercury":
		age = base / .2408467
	case "Venus":
		age = base / .61519726
	case "Mars":
		age = base / 1.8808158
	case "Jupiter":
		age = base / 11.862615
	case "Saturn":
		age = base / 29.447498
	case "Uranus":
		age = base / 84.016846
	case "Neptune":
		age = base / 164.79132
	}
	return age
}
