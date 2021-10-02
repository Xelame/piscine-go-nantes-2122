package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	number := 0
	if len(args) == 3 {
		if !(Atoi(args[0]) == 0 && args[0] != "0") || !(Atoi(args[2]) == 0 && args[2] != "0") {
			switch operator := args[1]; operator {
			case "+":
				number = Atoi(args[0]) + Atoi(args[2])
				break
			case "-":
				number = Atoi(args[0]) - Atoi(args[2])
				break
			case "/":
				number = Atoi(args[0]) / Atoi(args[2])
				break
			case "*":
				number = Atoi(args[0]) * Atoi(args[2])
				break
			case "%":
				number = Atoi(args[0]) % Atoi(args[2])
				break
			}
		}
		str := PrintNbr(number)
		if str != "(" && str != "-(" {
			os.Stdout.WriteString(str)
		}
	}
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

func PrintNbr(number int) string {
	Str := []rune{}
	if number < 0 {
		Str = append(Str, '-')
		number = -number
	}
	var reversedList []int
	for number >= 10 {
		reversedList = append(reversedList, number%10)
		number /= 10
	}
	reversedList = append(reversedList, number%10)
	for end := len(reversedList) - 1; end >= 0; end-- {
		Str = append(Str, rune(48+reversedList[end]))
	}
	Str = append(Str, '\n')
	return string(Str)
}
