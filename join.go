package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	} else {
		return Concat(strs[0], Concat(sep, Join(strs[1:], sep)))
	}
}
