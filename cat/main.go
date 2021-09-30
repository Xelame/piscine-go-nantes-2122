package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			panic(err)
		}
	} else {
		for _, arg := range os.Args[1:] {
			_, err := ioutil.ReadFile(arg)
			if err != nil {
				errorMessage := "ERROR: " + err.Error()
				for _, letter := range errorMessage {
					z01.PrintRune(letter)
				}
				z01.PrintRune('\n')
			}
			os.Exit(1)
		}
	}
}
