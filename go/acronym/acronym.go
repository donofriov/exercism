// Package acronym returns an acronym based on a long name
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate takes a long name string and returns a multi-letter acronym string
func Abbreviate(s string) string {
	// Split hypen words into separate words
	sanitizedString := strings.Replace(s, "-", " ", -1)

	// Removes all non-letters and non-whitespace
	sanitizedString = strings.Join(strings.Fields(strings.TrimSpace(sanitizedString)), " ")

	// Finds only letters and whitespace
	lettersRegEx, _ := regexp.Compile("[^a-zA-Z ]+")
	sanitizedString = lettersRegEx.ReplaceAllString(sanitizedString, "")

	// Capitalize only first letter of each word
	sanitizedString = strings.Title(sanitizedString)

	words := strings.Split(sanitizedString, " ")

	acronym := ""

	for _, word := range words {
		// Gets the first letter of each word
		acronym = acronym + string([]rune(word)[0])
	}

	return acronym
}
