package pangram

import "strings"

const testVersion = 1

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	runeMap := make(map[rune]bool)
	for _, runeVal := range input {
		runeMap[runeVal] = true
	}

	for _, runeVal := range "abcdefghijklmnopqrstuvwxyz" {
		if !runeMap[runeVal] {
			return false
		}
	}
	return true
}
