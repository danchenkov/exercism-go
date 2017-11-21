package hamming

import "errors"

const testVersion = 6

func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		return -1, errors.New("unequal lengths")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance += 1
		}
	}

	return distance, nil
}
