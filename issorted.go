package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	test := []int{}
	isOkay := false
	count := 0
	for i := 0; i < len(a)-1; i++ {
		test = append(test, f(a[i], a[i+1]))
	}
	for _, value := range test {
		if value == 0 {
			count++
			if count == len(test) {
				isOkay = true
			}
		}
	}
	return isOkay
}
