package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		c, e := ioutil.ReadAll(os.Stdin)
		if e == nil {
			printStr(string(c))
		} else {
			printStr(e.Error())
		}
	} else {
		for _, arg := range os.Args[1:] {
			b, err := ioutil.ReadFile(arg)
			if err != nil {
				errorMessage := "ERROR: " + err.Error()
				for _, letter := range errorMessage {
					z01.PrintRune(letter)
				}
				z01.PrintRune('\n')
				os.Exit(1)
			}
			printStr(string(b))
		}
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
