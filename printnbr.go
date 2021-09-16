package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(number int) {
	// Test if number is negative
	if number < 0 {
		z01.PrintRune('-')
		number -= 2 * number
	}
	// Create a loop to set up my power of 10
	for number > 10 {
		power := 1
		for number/power > 10 {
			power *= 10
		}
		// init some variable
		var count int32 = 1
		testing := power/10
		// surrounding our digit * 1^power
		for i := 1; i <= 9; i++ {
			if i*power < number && number < (i+1)*power {
				z01.PrintRune(48 + count)
				// test for the presence of one 0 or more
				for number%power < testing {
					z01.PrintRune(48)
					testing /= 10
				}
				// moving on to the digit in the right
				number %= power
			} else {
				count++
			}
		}
	}
	// Testing delayed for the last digit (unit)
	if number > 0 {
		count := rune(number % 10)
		z01.PrintRune(48+count)
	}
}
