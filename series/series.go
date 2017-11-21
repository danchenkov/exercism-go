package series

// fimport "strings"

const testVersion = 2

func All(n int, s string) []string {
	totalSubstrings := len(s) - n + 1
	allSubstrings := make([]string, totalSubstrings)

	if n > len(s) {
		return []string{}
	}

	for i := 0; i < totalSubstrings; i++ {
		allSubstrings[i] = s[i : n+i]
	}
	return allSubstrings
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		// return s + strings.Repeat(" ", n-len(s))
		return ""
	}
	return s[0:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[0:n], true
}
