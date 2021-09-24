package piscine

func SplitWhiteSpaces(s string) []string {
	arg := []string{}
	begin := 0
	for index, value := range s {
		if value == ' ' || value == '\n' || value == '\t' {
			arg = append(arg, s[begin:index])
			begin = index
		}
	}
	arg = append(arg, s[begin:])
	return arg
}
