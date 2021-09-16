package piscine

func StrLen(s string) int {
	count := 0
	for i, _ := range s {
		count = i
	}
	return count
}
