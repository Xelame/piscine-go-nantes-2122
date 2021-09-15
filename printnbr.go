package piscine

import "github.com/01-edu/z01"

func PrintNbr(nb int) {
	if nb < 0 {
		z01.PrintRune('-')
		nb -= 2 * nb
	}
	if nb > 100 {
		var count int32 = 1
		for i := 1; i <= 9; i++ {
			if i*100 < nb && nb < (i+1)*100 {
				z01.PrintRune(48 + count)
				nb -= i * 100
			} else {
				count++
			}
		}
		if nb > 10 {
			var count int32 = 1
			for i := 1; i <= 9; i++ {
				if i*10 < nb && nb < (i+1)*10 {
					z01.PrintRune(48 + count)
					nb -= i * 10
				} else {
					count++
				}
			}
		}
		if nb > 0 {
			var count int32 = 1
			for i := 1; i <= 9; i++ {
				if i == nb {
					z01.PrintRune(48 + count)
					nb -= i
				} else {
					count++
				}
			}
		}
	}
}
