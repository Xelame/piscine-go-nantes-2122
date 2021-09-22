package piscine

func ToUpper(aString string) string {
	listOfCharacter := []rune(aString)
	for _, value := range listOfCharacter {
		if 97 <= value && value <= 122 {
			value -= 32
		}
	}
	return string(listOfCharacter)
}
