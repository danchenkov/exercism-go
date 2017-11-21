package pangram

import "strings"

const testVersion = 1

func IsPangram(phrase string) bool {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	phrase = strings.ToLower(phrase)
	for _, char := range alphabet {
		if !strings.ContainsRune(phrase, char) {
			return false
		}
	}

	return true
}

// func IsPangram(input string) bool {
// 	alphabet := map[byte]int{
// 		'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0,
// 		'h': 0, 'i': 0, 'j': 0, 'k': 0, 'l': 0, 'm': 0, 'n': 0,
// 		'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0,
// 		'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0}

// 	for _, symbol := range alphabet {
// 		if symbol >= 'A' && symbol <= 'Z' {
// 			symbol = symbol - 'A' + 'a'
// 		}
// 		if symbol >= 'a' && symbol <= 'z' {
// 			alphabet[byte(symbol)] += 1
// 		}
// 	}

// 	for _, v := range alphabet {
// 		if v == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
