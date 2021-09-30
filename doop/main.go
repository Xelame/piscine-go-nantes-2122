package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[:1]
	number := 0
	if len(args) == 3 {
		switch operator := args[1]; operator {
		case "+":
			number = Atoi(args[0]) + Atoi(args[2])
		case "-":
			number = Atoi(args[0]) - Atoi(args[2])
		case "/":
			number = Atoi(args[0]) / Atoi(args[2])
		case "*":
			number = Atoi(args[0]) * Atoi(args[2])
		case "%":
			number = Atoi(args[0]) % Atoi(args[2])
		}
	}
	PrintNbr(number)
}

func Atoi(aString string) int {
	if aString != "" {
		powerOfTen := 1
		result := 0
		for index := len(aString) - 1; index >= 0; index-- {
			if 48 <= aString[index] && aString[index] <= 57 {
				result += powerOfTen * int(aString[index]%48)
				powerOfTen *= 10
			} else if (aString[index] == 45 || aString[index] == 43) && index == 0 {
				powerOfTen *= 10
			} else {
				return 0
			}
		}
		if aString[0] == 45 {
			result = -result
		}
		return result
	} else {
		return 0
	}
}

func PrintNbr(number int) {
	if number == -9223372036854775808 {
		number = 223372036854775808
		z01.PrintRune('-')
		z01.PrintRune('9')
	}
	if number < 0 {
		z01.PrintRune('-')
		number = -number
	}
	var reversedList []int
	for number >= 10 {
		reversedList = append(reversedList, number%10)
		number /= 10
	}
	reversedList = append(reversedList, number%10)
	for end := len(reversedList) - 1; end >= 0; end-- {
		z01.PrintRune(rune(48 + reversedList[end]))
	}
}
