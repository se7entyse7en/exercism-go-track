// Package bob implements the "Bob" exercise.
package bob

import (
	"regexp"
	"strings"
)

type condition struct {
	isQuestion, isYelling bool
}

var nonEmptyConditions = map[condition]string{
	condition{false, false}: "Whatever.",
	condition{false, true}:  "Whoa, chill out!",
	condition{true, false}:  "Sure.",
	condition{true, true}:   "Calm down, I know what I'm doing!",
}

// Hey reproduces a lackadaisical teenager with limited responses.
func Hey(remark string) string {
	strLen := len(strings.TrimSpace(remark))
	if strLen == 0 {
		return "Fine. Be that way!"
	}
	isQuestion := remark[strLen-1] == '?'
	isNotYelling, _ := regexp.MatchString(`([a-z]+)|(^[\W\d_]+$)`, remark)
	isYelling := !isNotYelling

	return nonEmptyConditions[condition{isQuestion, isYelling}]
}
