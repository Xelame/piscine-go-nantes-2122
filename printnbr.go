package piscine

import "github.com/01-edu/z01"

func PrintNbr(number int) {
	if number < 0 {
		z01.PrintRune('-')
		number -= 2 * number
	}
	for number > 10 {
		power := 1
		for number/power > 10 {
			power *= 10
		}
		var count int32 = 1
		testing := 1
		for i := 1; i <= 9; i++ {
			if i*power < number && number < (i+1)*power {
				z01.PrintRune(48 + count)
				for number%power < power/10*testing {
					z01.PrintRune(48)
					testing*=10
				}
				number %= power
			} else {
				count++
			}
		}
	}
	if number > 0 {
		var count int32 = 1
		for i := 1; i <= 9; i++ {
			if i == number {
				z01.PrintRune(48 + count)
			} else {
				count++
			}
		}
	}
}
