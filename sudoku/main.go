package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	for _, i := range arg {
		fmt.Println(i)
	}
}
