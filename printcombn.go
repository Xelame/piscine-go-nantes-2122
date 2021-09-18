package piscine

import "github.com/01-edu/z01"

func PrintCombN(number int) {
	table := [9]int{}
	isAscencingOrder := true
	isFirstOutput := true
	// Looping while the table is different from [9, ?, ?, ...]
	for table[0] <= 9 {
		isAscencingOrder = true
		// Test if PrintCombN(number > 1) else do out of range
		if number > 1 {
			// Do counting from 0 to 999999999 maximum
			for test := number; test > 1; test-- {
				if table[test-1]%10 == 0 && table[test-1] != 0 {
					table[test-2]++
					table[test-1] %= 10
					table[test-1] += table[test-2]
				}
			}
			// Do tests if digits is different and in ascencing order in the table
			for lower := number; lower > 1; lower-- {
				if !(table[lower-1] > table[lower-2] && table[lower-1] != table[lower-2]) {
					isAscencingOrder = false
				}
			}
		}
		// if all test works
		if isAscencingOrder {
			// if the first output don't write a comma before
			if !isFirstOutput {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			isFirstOutput = false
			// Writing number
			for index := 0; index < len(table[:number]); index++ {
				z01.PrintRune(rune(48 + table[index]))
			}
		}
		// Indentation for counting
		table[number-1]++
	}
	z01.PrintRune('\n')
}
