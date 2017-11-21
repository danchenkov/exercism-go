package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	runes := []rune(strings.ToLower(word))
	verification := map[rune]struct{}{}
	for _, letter := range runes {
		if unicode.IsLetter(letter) {
			if _, ok := verification[letter]; ok {
				return false
			}
			verification[letter] = struct{}{}
		}
	}
	return true
}
