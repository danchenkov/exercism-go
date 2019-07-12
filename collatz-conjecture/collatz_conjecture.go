package collatzconjecture

import "errors"

func CollatzConjecture(input int) (int, error) {
	counter := 0
	switch {
	case input < 0:
		return 0, errors.New("negative input")
	case input == 0:
		return 0, errors.New("zero input")
	case input == 1:
		return 0, nil
	default:
		for ; input != 1; counter++ {
			if input%2 == 0 {
				input = input / 2
			} else {
				input = input*3 + 1
			}
		}
		return counter, nil
	}
}
