package two_pointers

import (
	"regexp"
	"strings"
)

func is_palindrome(s string) bool {
	specialCharactersRegex := regexp.MustCompile(`[\p{P}\p{S}]+`)
	lowerCasedString := strings.ToLower(specialCharactersRegex.ReplaceAllString(s, ""))
	cleaned := strings.ReplaceAll(lowerCasedString, " ", "")

	length := len(cleaned)

	for i := 0; i < length; i++ {
		if cleaned[i] != cleaned[length-i-1] {
			return false
		}
	}

	return true
}
