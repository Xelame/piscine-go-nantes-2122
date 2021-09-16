package piscine

func StrRev(s string) string {
	sChangeable := []byte(s)
	var temp byte
	for beginning := 0; beginning < len(sChangeable); beginning++ {
		if beginning < len(sChangeable)-beginning-1 {
			temp = sChangeable[beginning]
			sChangeable[beginning] = sChangeable[len(sChangeable)-beginning-1]
			sChangeable[len(sChangeable)-beginning-1] = temp
		}
	}
	return string(sChangeable)
}
