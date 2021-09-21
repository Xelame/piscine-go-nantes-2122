package piscine

func Fibonacci(index int) int {
	if index <= 2 {
		return 1
	} else {
		return Fibonacci(index-1) + Fibonacci(index-2)
	}
}
