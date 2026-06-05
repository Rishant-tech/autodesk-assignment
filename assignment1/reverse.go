package main

import "unicode"

// ReverseWords reverses each alphanumeric word in the string while preserving
// the original order of words and all non-word characters.
func ReverseWords(s string) string {
	runes := []rune(s)
	out := make([]rune, 0, len(runes))

	for i := 0; i < len(runes); {
		if !isWordRune(runes[i]) {
			out = append(out, runes[i])
			i++
			continue
		}

		start := i
		for i < len(runes) && isWordRune(runes[i]) {
			i++
		}

		for j := i - 1; j >= start; j-- {
			out = append(out, runes[j])
		}
	}

	return string(out)
}

func isWordRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}
