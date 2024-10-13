package stringtocamelcase

import (
	"regexp"
	"strings"
	"unicode"
)

// ToCamelCase converts dash/underscore delimited words into camel casing.
// The first word within the output is capitalized only if the original word was capitalized (known as Upper Camel Case, also often referred to as Pascal case).
// The next words are always capitalized.
//
// Examples:
//   - "the-stealth-warrior" gets converted to "theStealthWarrior"
//   - "The_Stealth_Warrior" gets converted to "TheStealthWarrior"
//   - "The_Stealth-Warrior" gets converted to "TheStealthWarrior"
func ToCamelCase(s string) string {
	wordsSlice := regexp.MustCompile(`[-_]`).Split(s, -1)
	var result string

	for iWord, word := range wordsSlice {
		for iLetter, letter := range word {
			if iWord != 0 && iLetter == 0 && unicode.IsLower(letter) {
				result += strings.ToUpper(string(letter))
				continue
			}
			result += string(letter)
		}
	}

	return result
}
