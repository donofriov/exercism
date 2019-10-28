// Package twofer prints out One for you and one for me where
// you is possibly replaced with a name
package twofer

import "fmt"

// ShareWith returns a string with either a persons name
// or a hardcoded 'you' catchall when name is not defined.
func ShareWith(name string) string {
	if (name == "") {
		name = "you"
	}

	return fmt.Sprintf("One for %v, one for me.", name)
}
