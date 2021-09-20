package piscine

func StrRev(s string) string {
	runes := []rune(s)
	for beginning, end := 0, len(runes)-1; beginning < end; beginning, end = beginning+1, end-1 {
		runes[beginning], runes[end] = runes[end], runes[beginning]
	}
	return string(runes)
}
