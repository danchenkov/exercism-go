package reverse

const testVersion = 1

func String(input string) string {
	runes := []rune(input)
	size := len(runes)
	output := make([]rune, size)

	for index, rune := range runes {
		output[size-1-index] = rune
	}
	return string(output)
}

// Slower:
// func String(input string) string {
// 	runes := []rune(input)

// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }
