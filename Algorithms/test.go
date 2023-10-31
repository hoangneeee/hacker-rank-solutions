package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "saveChangesInTheEditor"
		var words []string
		wordStart := 0

		for i, char := range s {
			if i > 0 && unicode.IsUpper(char) {
				words = append(words, s[wordStart:i])
				wordStart = i
			}
		}

		if wordStart < len(s) {
			words = append(words, s[wordStart:])
		}
		fmt.Printf("%v", len(words))
}
