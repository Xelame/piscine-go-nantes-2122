package piscine

func NRune(aString string, index int) rune {
	var result rune
	for i, letter := range aString {
		if i == index-1 {
			result = rune(letter)
		}
	}
	return result
}
