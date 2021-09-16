package piscine

import "github.com/01-edu/z01"

func PrintNbr(number int) {
	// Test if number is negative
	if number < 0 {
		z01.PrintRune('-')
		number = -number
	}
	var reversedNumberInList []int
	for number >= 10 {
		reversedNumberInList = append(reversedNumberInList, number%10)
		number /= 10
	}
	reversedNumberInList = append(reversedNumberInList, number%10)
	for end := len(reversedNumberInList) - 1; end >= 0; end-- {
		z01.PrintRune(rune(48 + reversedNumberInList[end]))
	}
}
