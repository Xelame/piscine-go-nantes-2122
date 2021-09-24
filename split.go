package piscine

func Split(s, sep string) []string {
	arg := []string{}
	begin := -1
	for index, value := range s {
		if value == rune(sep[0]) && s[index:index+len(sep)] == sep {
			if s[begin+1:index] != "" {
				if begin != -1 {
					arg = append(arg, s[begin+len(sep):index])
				} else {
					arg = append(arg, s[begin+1:index])
				}
			}
			begin = index
		}
	}
	arg = append(arg, s[begin+len(sep):])
	return arg
}
