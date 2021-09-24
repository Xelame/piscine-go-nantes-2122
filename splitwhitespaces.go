package piscine

func SplitWhiteSpaces(s string) []string {
	arg := []string{}
	begin := -1
	for index, value := range s {
		if value == ' ' || value == '\n' || value == '\t' {
			if s[begin+1:index] != "" {
				arg = append(arg, s[begin+1:index])
			}
			begin = index
		}
	}
	arg = append(arg, s[begin+1:])
	return arg
}
