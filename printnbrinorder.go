package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(number int) {
	// Test if number is negative
	/*
		if number < 0 {
			z01.PrintRune('-')
		}
	*/
	var NumberInList []int
	for number >= 10 {
		NumberInList = append(NumberInList, number%10)
		number /= 10
	}
	NumberInList = append(NumberInList, number%10)
	SortIntegerTable(NumberInList)
	for end := range NumberInList {
		z01.PrintRune(rune(48 + NumberInList[end]))
	}
}
