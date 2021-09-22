package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(number int) {
	// Test if number is negative
	/*
		if number < 0 {
			z01.PrintRune('-')
		}
	*/
	var reversedNumberInList []int
	for number >= 10 || number <= -10 {
		reversedNumberInList = append(reversedNumberInList, number%10)
		number /= 10
	}
	reversedNumberInList = append(reversedNumberInList, number%10)
	SortIntegerTable(reversedNumberInList)
	for end := range reversedNumberInList {
		if reversedNumberInList[end] > 0 {
			z01.PrintRune(rune(48 + reversedNumberInList[end]))
		} else {
			z01.PrintRune(rune(48 - reversedNumberInList[end]))
		}
	}
}
