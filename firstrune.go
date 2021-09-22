package piscine

func FirstRune(aString string) rune {
	var result rune
	for i := range aString {
		if i == 0 {
			result = rune(aString[i])
		}
	}
	return result
}
