package piscine

func Map(f func(int) bool, a []int) []bool {
	b := []bool{}
	for index := range a {
		b = append(b, f(a[index]))
	}
	return b
}
