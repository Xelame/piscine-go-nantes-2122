package piscine

import "github.com/01-edu/z01"

func PrintNbr(nb int) {
	if nb < 0 {
		z01.PrintRune('-')
		nb -= 2 * nb
	}
	for nb > 10 {
		powerNumber := 1
		limits := 0
		for powerExp := 0; powerExp <= limits; powerExp++ {
			powerNumber = powerNumber * 10
			if 10*powerNumber < nb && nb < 90*powerNumber {
				limits++
			}
		}
		if 9*powerNumber > nb && nb > powerNumber {
			var count int32 = 1
			for i := 1; i <= 9; i++ {
				if i*powerNumber < nb && nb < (i+1)*powerNumber {
					z01.PrintRune(48 + count)
					nb %= powerNumber

				} else {
					count++
				}
			}
		}
	}
	if nb > 0 {
		var count int32 = 1
		for i := 1; i <= 9; i++ {
			if i == nb {
				z01.PrintRune(48 + count)
			} else {
				count++
			}
		}
	}

}
