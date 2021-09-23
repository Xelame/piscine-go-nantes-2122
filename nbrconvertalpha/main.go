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
		if IsAlpha(rune(value + int(converter))) {
			z01.PrintRune(rune(value + int(converter)))
		} else {
			z01.PrintRune(rune(32))
		}
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

func IsAlpha(aRune rune) bool {
	return (65 <= int(aRune) && int(aRune) <= 90) || (97 <= int(aRune) && int(aRune) <= 122)
}
