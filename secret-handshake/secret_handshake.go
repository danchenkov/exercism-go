package secret

const testVersion = 2

func Handshake(code uint) (decodedArray []string) {
	var secretMessages = []string{
		"wink",
		"double blink",
		"close your eyes",
		"jump",
	}
	var i uint
	for i = 0; i < 4; i++ {
		if code&(1<<i) != 0 {
			decodedArray = append(decodedArray, secretMessages[i])
		}
	}
	if code&(1<<4) != 0 {
		return reverseArray(decodedArray)
	}
	return decodedArray
}

func reverseArray(input []string) []string {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}
