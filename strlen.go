package piscine

func StrLen(s string) int {
	count := 0
	for index, value := range s {
		count++
		if value > 127 {
			count++
		}
	}
	return count
}
