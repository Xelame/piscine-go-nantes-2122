package piscine

import "fmt"

func PointOne(n *int) {
	*n += 1
	fmt.Println(*n)
}
