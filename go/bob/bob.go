// Package bob should have a package comment that summarizes what it"s about.
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	lowerRegEx, _ := regexp.Compile("[a-z]")
	lettersRegEx, _ := regexp.Compile("[a-zA-Z]")
	letters := lettersRegEx.MatchString(remark)
	yell := !lowerRegEx.MatchString(remark)
	question := strings.HasSuffix(remark, "?")

	if remark == "" {
		return "Fine. Be that way!"
	} else if !letters && question {
		return "Sure."
	} else if !letters {
		return "Whatever."
	} else if yell && question {
		return "Calm down, I know what I'm doing!"
	} else if question {
		return "Sure."
	} else if yell {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}
}
