package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fichier := os.Args
	SortTable(fichier)
	for _, arg := range fichier[1:] {
		begin := 0
		isFirst := true
		for i := len(arg) - 1; i >= 0; i-- {
			if arg[i] == '/' && isFirst {
				begin = i - 1
				isFirst = !isFirst
			}
		}
		for i, value := range arg {
			if i >= begin {
				z01.PrintRune(value)
			}
		}
		z01.PrintRune('\n')
	}
}

func SortTable(table []string) {
	for min_index := 0; min_index < len(table)-1; min_index++ {
		for index := min_index + 1; index < len(table); index++ {
			if Sum(table[index]) < Sum(table[min_index]) {
				Swap(&table[index], &table[min_index])
			}
		}
	}
}

func Swap(a *string, b *string) {
	*a, *b = *b, *a
}

func Sum(table string) int {
	total := 0
	for _, value := range table {
		total += int(value)
	}
	return total
}
