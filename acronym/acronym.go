package acronym

import (
	"strings"
)

const testVersion = 3

func Abbreviate(input string) (output string) {
	o := strings.Split(strings.Replace(input, "-", " ", -1), " ")
	for i := 0; i < len(o); i++ {
		output += strings.ToUpper(o[i][0:1])
	}
	return
}
