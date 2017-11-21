package raindrops

import "strconv"

const testVersion = 3

func Convert(i int) (s string) {
	if (i % 3) == 0 {
		s = "Pling"
	}
	if (i % 5) == 0 {
		s = s + "Plang"
	}
	if (i % 7) == 0 {
		s = s + "Plong"
	}
	if s == "" {
		return strconv.Itoa(i)
	}
	return s
}
