package piscine

func ToUpper(aString string) string {
	listOfCharacter := []rune(aString)
	for _, value := range listOfCharacter {
		if 65 <= value && value <= 90 {
			value += 32
		}
	}
	return string(listOfCharacter)
}
