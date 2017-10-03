package hamming

import "errors"

const testVersion = 6

func Distance(a, b string) (int, error) {
	var distance int
	if len(a) != len(b) {
		return 0, errors.New("Strings must be equal")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance += 1
		}
	}
	return distance, nil
}
