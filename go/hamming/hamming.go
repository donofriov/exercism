package hamming

import "errors"

// Distance function taking in two 'sequence' strings and
// returns the total difference of characters as an int
//
// Throws an error if the sequences aren't the same length
func Distance(a, b string) (int, error) {
	count := 0
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count++
			}
		}
		return count, nil
	}

	return 0, errors.New("Sequences are different lengths")
}
