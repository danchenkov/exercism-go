package bob

import "strings"

const testVersion = 3

func Hey(phrase string) string {
	phrase = strings.TrimSpace(phrase)
	if strings.ToUpper(phrase) == phrase && strings.ContainsAny(phrase, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return "Whoa, chill out!"
	}
	if strings.HasSuffix(phrase, "?") {
		return "Sure."
	}
	if phrase == "" {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
