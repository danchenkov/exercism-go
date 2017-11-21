package accumulate

const testVersion = 1

func Accumulate(input []string, f func(string) string) []string {
	output := make([]string, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = f(input[i])
	}
	return output
}
