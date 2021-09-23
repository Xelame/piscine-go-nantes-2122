package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(number int) {
	var NumberInList []int
	for number >= 10 {
		NumberInList = append(NumberInList, number%10)
		number /= 10
	}
	NumberInList = append(NumberInList, number%10)
	SortIntegerTable(NumberInList)
	for _, value := range NumberInList {
		z01.PrintRune(rune(48 + value))
	}
}
