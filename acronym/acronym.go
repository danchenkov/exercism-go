package acronym

import (
	"strings"
)

const testVersion = 3

func Abbreviate(input string) (output string) {
	words := strings.Split(strings.Replace(input, "-", " ", -1), " ")
	for _, word := range words {
		if len(string(word)) > 0 {
			output += strings.ToUpper(string(word[0]))
		}
	}
	return output
}

// Regexp is slower

// func Abbreviate(input string) (output string) {
// 	words := regexp.MustCompile("(?: |-)+").Split(input, -1)
// 	for _, word := range words {
// 		output += strings.ToUpper(string(word[0]))
// 	}
// 	return output
// }
