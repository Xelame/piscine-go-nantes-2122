package piscine

func LastRune(aString string) rune {
	var result rune
	for _, letter := range aString {
		result = rune(letter)
	}
	return result
}
