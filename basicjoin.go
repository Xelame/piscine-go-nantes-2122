package piscine

func BasicJoin(elems []string) string {
	if len(elems) == 0 {
		return ""
	}
	if len(elems) == 1 {
		return elems[0]
	} else {
		return Concat(elems[0], BasicJoin(elems[1:]))
	}
}
