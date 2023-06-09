package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	result := make([]int, max-min)
	for i := 0; i < len(result); i++ {
		result[i] = min + i
	}
	return result
}
