package piscine

func StrLen(s string) int {
	count := 0
	for i := range s {
		count = i
	}
	return count + 1
}
