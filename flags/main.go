package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fichier := os.Args[1:]
	if len(fichier) == 0 || fichier[0] == "--help" || fichier[0] == "-h" {
		fmt.Println("--insert\n  -i\n	 This flag inserts the string into the string passed as argument.\n--order\n  -o\n	 This flag will behave like a boolean, if it is called it will order the argument.")
	} else {
		for i := 0; i < len(fichier); i++ {
			if len(fichier[i]) > 9 && fichier[i][:9] == "--insert=" {
				fichier[len(fichier)-1] = Concat(fichier[len(fichier)-1], fichier[i][9:])
			} else if len(fichier[i]) > 3 && fichier[i][:3] == "-i=" {
				fichier[len(fichier)-1] = Concat(fichier[len(fichier)-1], fichier[i][3:])
			}
			if fichier[i] == "--order" || fichier[i] == "-o" {
				for v := 0; v < len(fichier); v++ {
					if !(fichier[v] == "--order" || fichier[v] == "-o" || (len(fichier[v]) > 9 && fichier[v][:9] == "--insert=") || (len(fichier[v]) > 3 && fichier[v][:3] == "-i=")) {
						temp := []rune(fichier[i+1])
						SortTable(temp)
						fichier[i+1] = string(temp)
					}
				}
			}
		}
		for _, arg := range fichier {
			if !(arg == "--order" || arg == "-o") {
				for _, value := range arg {
					z01.PrintRune(value)
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func SortTable(table []rune) {
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

func Concat(s1 string, s2 string) string {
	return s1 + s2
}
