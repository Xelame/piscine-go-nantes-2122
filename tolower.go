package piscine

func ToLower(aString string) string {
	listOfCharacter := []rune(aString)
	for index, value := range listOfCharacter {
		if 65 <= value && value <= 90 {
			listOfCharacter[index] += 32
		}
	}
	return string(listOfCharacter)
}
