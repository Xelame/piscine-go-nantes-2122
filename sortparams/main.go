package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fichier := os.Args
	SortIntegerTable([]rune(fichier[0]))
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

func SortIntegerTable(table []rune) {
	for min_index := 0; min_index < len(table)-1; min_index++ {
		for index := min_index + 1; index < len(table); index++ {
			if table[index] < table[min_index] {
				Swap(&table[index], &table[min_index])
			}
		}
	}
}

func Swap(a *rune, b *rune) {
	*a, *b = *b, *a
}
