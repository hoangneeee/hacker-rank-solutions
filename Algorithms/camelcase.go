func camelcase(s string) int32 {
	// Write your code here
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
	return int32(len(words))
}