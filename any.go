package piscine

func Any(f func(string) bool, a []string) bool {
	isTrue := false
	for _, value := range a {
		if !isTrue {
			isTrue = f(value)
		}
	}
	return isTrue
}
