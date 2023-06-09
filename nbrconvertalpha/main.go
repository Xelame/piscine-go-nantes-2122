package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	haveSomething := false
	fichier := os.Args
	listOfInt := []int{}
	var converter rune = 96
	for _, arg := range fichier[1:] {
		if arg == "--upper" {
			converter = 64
		} else if IsNumeric(arg) && 0 < BasicAtoi(arg) && BasicAtoi(arg) < 27 {
			listOfInt = append(listOfInt, BasicAtoi(arg))
		} else {
			listOfInt = append(listOfInt, ' ')
		}
	}
	for _, value := range listOfInt {
		if value != 32 {
			z01.PrintRune(rune(value + int(converter)))
		} else {
			z01.PrintRune(rune(32))
		}
		haveSomething = true
	}
	if haveSomething {
		z01.PrintRune('\n')
	}
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

func IsNumeric(aString string) bool {
	isNum := true
	for _, value := range aString {
		if !(48 <= int(value) && int(value) <= 57) && isNum {
			isNum = !isNum
		}
	}
	return isNum
}
