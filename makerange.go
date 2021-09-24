package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	result := []int{max - min}
	for i := 0; i < len(result); i++ {
		result[i] = min + 1
	}
	return result
}
