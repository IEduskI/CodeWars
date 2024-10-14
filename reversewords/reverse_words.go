package reversewords

import "strings"

// ReverseWords
// Complete the function that accepts a string parameter, and reverses each word in the string. All spaces in the string should be retained.
//
// Examples
// "This is an example!" ==> "sihT si na !elpmaxe"
// "double  spaces"      ==> "elbuod  secaps"
func ReverseWords(str string) string {
	var result string
	words := strings.Split(str, " ")

	for indexW, word := range words {
		i := 1
		for i <= len(word) {
			result += string(word[len(word)-i])
			i++
		}
		if indexW+1 != len(words) {
			result += " "
		}
	}
	return result
}
