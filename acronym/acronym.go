package acronym

import "strings"

const testVersion = 2

func Abbreviate(readableText string) string {
	var fancyAcronym string
	const uppers string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	splits := strings.FieldsFunc(readableText, func(r rune) bool {
		switch r {
		case ' ', '-':
			return true
		}
		return false
	})

	var str string
	for i := 0; i < len(splits); i++ {
		str = splits[i]
		capitalIndices := make(map[int]bool)
		for j := 0; j < len(str); j++ {
			if strings.ContainsAny(str[j:j+1], uppers) || j == 0 {
				capitalIndices[j] = true
				if !capitalIndices[j-1] {
					fancyAcronym += str[j : j+1]
				}
			} else {
				capitalIndices[j] = false
			}
		}
	}
	return strings.ToUpper(fancyAcronym)
}
