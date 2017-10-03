package bob

import (
	"regexp"
	"strings"
)

const testVersion = 3

func Hey(input string) string {
	if Shouting(input) {
		return "Whoa, chill out!"
	}
	if Question(input) {
		return "Sure."
	}
	if Empty(input) {
		return "Fine. Be that way!"
	}
	return "Whatever."
}

func Empty(input string) bool {
	result, _ := regexp.Match(`^\s*$`, []byte(input))
	return result
}

func Question(input string) bool {
	result, _ := regexp.Match(`.*\?\s*$`, []byte(input))
	return result
}

func Shouting(input string) bool {
	return strings.ToUpper(input) == input && strings.ContainsAny(input, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
