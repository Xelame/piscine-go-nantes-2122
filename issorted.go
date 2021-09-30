package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	test := []int{}
	isOkay := false
	count1 := 0
	count2 := 0
	for i := 0; i < len(a)-1; i++ {
		test = append(test, f(a[i], a[i+1]))
	}
	for _, value := range test {
		if value == 0 {
			count1++
			if count1 == len(test) {
				isOkay = true
			}
		}
		if value > 0 {
			count2++
			if count2 == len(test) {
				isOkay = true
			}
		}
	}
	return isOkay
}
