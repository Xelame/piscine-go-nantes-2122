package piscine

func FirstRune(aString string) rune {
	var result rune
	for i, letter := range aString {
		if i == 0 {
			result = rune(letter)
		}
	}
	return result
}
