package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fichier := os.Args
	listOfInt := []int{}
	var converter rune = 96
	for _, arg := range fichier[1:] {
		if arg == "--upper" {
			converter = 64
		} else {
			listOfInt = append(listOfInt, BasicAtoi(arg))
		}
	}
	for _, value := range listOfInt {
		z01.PrintRune(rune(value + int(converter)))

	}
	z01.PrintRune('\n')
}

func BasicAtoi(aString string) int {
	powerOfTen := 1
	result := 0
	for index := range aString {
		result += powerOfTen * int(aString[len(aString)-1-index]%48)
		powerOfTen *= 10
	}
	return result
}
