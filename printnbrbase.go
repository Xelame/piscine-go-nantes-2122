package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(number int, base string) {
	isOkay := true
	listOfNumberInMyBase := []rune(base)
	for i := len(listOfNumberInMyBase) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if listOfNumberInMyBase[i] == listOfNumberInMyBase[j] && isOkay {
				z01.PrintRune('N')
				z01.PrintRune('V')
				isOkay = false
			}
		}
	}
	if isOkay {
		if number < 0 {
			z01.PrintRune('-')
		}
		var reversedNumberInList []int
		for number >= len(base) || number <= -len(base) {
			reversedNumberInList = append(reversedNumberInList, number%len(base))
			number /= len(base)
		}
		reversedNumberInList = append(reversedNumberInList, number%len(base))
		for end := len(reversedNumberInList) - 1; end >= 0; end-- {
			if reversedNumberInList[end] > 0 {
				z01.PrintRune(listOfNumberInMyBase[reversedNumberInList[end]%48])
			} else {
				z01.PrintRune(listOfNumberInMyBase[-(reversedNumberInList[end] % 48)])
			}
		}
	}
}
