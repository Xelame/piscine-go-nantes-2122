package piscine

import "github.com/01-edu/z01"

func PrintCombN(number int) {
	table := [9]int{}
	isAscencingOrder := true
	isFirst := true
	for table[0] <= 9 {
		isAscencingOrder = true
		if number > 1 {
			for test := number; test > 1; test-- {
				if table[test-1]%10 == 0 && table[test-1] != 0 {
					table[test-2]++
					table[test-1] %= 10
				}
			}
			for lower := number; lower > 1; lower-- {
				if !(table[lower-1] > table[lower-2] && table[lower-1] != table[lower-2]) {
					isAscencingOrder = false
				}
			}
		}
		if isAscencingOrder {
			if !isFirst {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			isFirst = false
			for index := 0; index < len(table[:number]); index++ {
				z01.PrintRune(rune(48 + table[index]))
			}
		}
		table[number-1]++
	}
	z01.PrintRune('\n')
}
