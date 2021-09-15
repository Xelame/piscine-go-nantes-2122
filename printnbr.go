package piscine

import "github.com/01-edu/z01"

func PrintNbr(nb int) {
	if nb < 0 {
		z01.PrintRune('-')
		nb -= 2 * nb
	}
	for nb > 10 {
		power := 1
		for nb/power > 10 {
			power *= 10
		}
		if nb > power {
			var count int32 = 1
			for i := 1; i <= 9; i++ {
				if i*power < nb && nb < (i+1)*power {
					z01.PrintRune(48 + count)
					nb %= power
				} else {
					count++
				}
			}
		}
	}
	if nb < 9 {
		var count int32 = 0
		for i := 0; i <= 9; i++ {
			if i == nb {
				z01.PrintRune(48 + count)
			} else {
				count++
			}
		}
	}
}
