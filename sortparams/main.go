package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fichier := os.Args[1:]
	SortTable(fichier)
	for _, arg := range fichier {
		for _, value := range arg {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}

func SortTable(table []string) {
	for min_index := 0; min_index < len(table)-1; min_index++ {
		for index := min_index + 1; index < len(table); index++ {
			if table[index] < table[min_index] {
				Swap(&table[index], &table[min_index])
			}
		}
	}
}

func Swap(a *string, b *string) {
	*a, *b = *b, *a
}
