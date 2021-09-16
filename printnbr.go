package piscine

import "github.com/01-edu/z01"

func PrintNbr(number int) {
	// if is a negative number
	if number < 0 {
		z01.PrintRune('-')
		number -= 2 * number
	}
	// search the max power of 10 to read the first digit of number 
	for number > 10 {
		power := 1
		for number / power > 10 {
			power *= 10
		}
		// search which digits is bitween 1 and 9
		if number > power {
			var count int32 = 1
			for i := 1; i <= 9; i++ {
				if i*power < number && number < (i+1)*power {
					z01.PrintRune(48 + count)
					number %= power
				} else {
					count++
				}
			}
		}
	}
	// Same thing just for the last digit
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
